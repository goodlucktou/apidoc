// SPDX-License-Identifier: MIT

// Package input 用于处理输入的文件，从代码中提取基本的注释内容。
//
// 多行注释和单行注释在处理上会有一定区别：
//  - 单行注释，风格相同且相邻的注释会被合并成一个注释块；
//  - 单行注释，风格不相同且相邻的注释会被按注释风格多个注释块；
//  - 多行注释，即使两个注释释块相邻也会被分成两个注释块来处理。
package input

import (
	"bytes"
	"context"
	"io/ioutil"
	"strconv"
	"sync"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"

	"github.com/caixw/apidoc/v5/internal/lang"
	"github.com/caixw/apidoc/v5/internal/locale"
	"github.com/caixw/apidoc/v5/message"
	opt "github.com/caixw/apidoc/v5/options"
)

// Block 解析出来的注释块
type Block struct {
	File string
	Line int
	Data []byte
}

// Parse 分析源代码，获取相应的文档内容。
//
// 当所有的代码块已经放入 Block 之后，Block 会被关闭。
//
// 所有与解析有关的错误均通过 h 输出；
// 而 inputs 中的相关错误，会通过返回的 SyntaxError 表示。
func Parse(ctx context.Context, h *message.Handler, inputs ...*opt.Input) (chan Block, *message.SyntaxError) {
	if len(inputs) == 0 {
		return nil, message.NewError("", "inputs", 0, locale.ErrRequired)
	}

	opts := make([]*options, 0, len(inputs))
	for index, item := range inputs {
		field := "inputs[" + strconv.Itoa(index) + "]"
		if item == nil {
			return nil, message.NewError("", field, 0, locale.ErrRequired)
		}
		opt, err := buildOptions(item)
		if err != nil {
			err.Field = field + "." + err.Field
			return nil, err
		}

		opts = append(opts, opt)
	}

	data := make(chan Block, 500)

	go func() {
		wg := &sync.WaitGroup{}
		for _, opt := range opts {
			select {
			case <-ctx.Done():
				return
			default:
				parse(ctx, data, h, wg, opt)
			}
		}
		wg.Wait()

		close(data)
	}()

	return data, nil
}

func parse(ctx context.Context, data chan Block, h *message.Handler, wg *sync.WaitGroup, o *options) {
	for _, path := range o.paths {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func(path string) {
				parseFile(data, h, path, o)
				wg.Done()
			}(path)
		}
	} // end for
}

// 分析 path 指向的文件。
//
// NOTE: parseFile 内部不能有协程处理代码。
func parseFile(channel chan Block, h *message.Handler, path string, o *options) {
	data, err := readFile(path, o.encoding)
	if err != nil {
		h.Error(message.Erro, message.WithError(path, "", 0, err))
		return
	}

	ret := lang.Parse(path, data, o.blocks, h)
	for line, data := range ret {
		channel <- Block{
			File: path,
			Line: line,
			Data: data,
		}
	}
}

// 以指定的编码方式读取内容。
func readFile(path string, encoding encoding.Encoding) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if encoding == nil {
		return data, nil
	}

	reader := transform.NewReader(bytes.NewReader(data), encoding.NewDecoder())
	return ioutil.ReadAll(reader)
}
