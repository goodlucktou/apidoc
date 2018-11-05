// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locale

import (
	"golang.org/x/text/language"
)

// 各个语种的语言对照表，通过相应文件的 init() 作初始化这样
// 在删除文件是，就自动删除相应的语言文件，不需要手修改代码。
var locales = map[language.Tag]map[string]string{}

// 各个语言需要翻译的所有字符串
const (
	// 与 flag 包相关的处理
	FlagUsage = `%v 是一个 RESTful API 文档生成工具。

参数：
%v

源代码采用 MIT 开源许可证，发布于 %v
详细信息可访问官网 %v
`
	FlagHUsage              = "显示帮助信息"
	FlagVUsage              = "显示版本信息"
	FlagLanguagesUsage      = "显示所有支持的语言"
	FlagGUsage              = "创建一个默认的配置文件"
	FlagEncodingsUsage      = "显示支持的编码方式"
	FlagWDUsage             = "指定工作目录，默认为当前目录"
	FlagPprofUsage          = "指定一种调试输出类型，可以为 %s 或是 %s"
	FlagVersionBuildWith    = "%v %v build with %v\n"
	FlagVersionCommitHash   = "commit hash %v\n"
	FlagSupportedLanguages  = "目前支持以下语言 %v\n"
	FlagSupportedEncodings  = "目前支持以下编码 https://www.iana.org/assignments/character-sets/character-sets.xhtml\n"
	FlagConfigWritedSuccess = "配置内容成功写入 %v"
	FlagPprofWritedSuccess  = "pprof 的相关数据已经写入到 %v"
	FlagInvalidPprrof       = "无效的 pprof 参数"

	VersionInCompatible = "当前程序与配置文件中指定的版本号不兼容"
	Complete            = "完成！文档保存在：%v，总用时：%v"

	// 错误信息，可能在地方用到
	ErrRequired              = "不能为空"
	ErrMustEmpty             = "只能为空"
	ErrInvalidFormat         = "格式不正确"
	ErrDirNotExists          = "目录不存在"
	ErrUnsupportedInputLang  = "无效的输入语言：%v"
	ErrNotFoundEndFlag       = "找不到结束符号"
	ErrNotFoundSupportedLang = "该目录下没有支持的语言文件"
	ErrUnknownTag            = "不认识的标签"
	ErrDuplicateTag          = "重复的标签"
	ErrUnsupportedEncoding   = "不支持的编码方式：%v"
	ErrDirIsEmpty            = "目录下没有需要解析的文件"
	ErrInvalidValue          = "无效的值"
	ErrInvalidOpenapi        = "openapi 内容错误：字段：%s；错误内容：%s"
	ErrDuplicateRoute        = "重复的路由项 %s:%s"
	ErrPathNotMatchParams    = "地址参数不匹配"
	ErrPathInvalid           = "地址格式错误"
	ErrDuplicateValue        = "重复的值"
	ErrMessage               = "%s 位于 %s:%d 的 %s"
	ErrMessageWithError      = "%s[%s] 位于 %s:%d 的 %s"
	//ErrSyntax                = "在[%s:%d]出现语法错误[%s]"
	//ErrConfig                = "配置文件[%s]中配置项[%s]错误[%s]"
	//ErrOptions       = "配置项[%s]错误[%s]"
	ErrApidocExists  = "相同组名的 @apidoc 标签已经存在"
	ErrInvalidMethod = "无效的请求方法"
	ErrMethodExists  = "该请求方法已经存在"
	ErrInvalidTag    = "无效的标签 %s"

	// logs
	InfoPrefix  = "[INFO] "
	WarnPrefix  = "[WARN] "
	ErrorPrefix = "[ERRO] "
)
