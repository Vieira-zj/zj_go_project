package control

import (
	"text/template"
)

var resFuncMap = make(template.FuncMap)

// GetFuncMap 返回funcmap
func GetFuncMap() template.FuncMap {
	return resFuncMap
}

// RegisterFuncMap 初始化funcmap
func RegisterFuncMap() {
	resFuncMap["hello"] = hello
	resFuncMap["helloMsg"] = helloMsg
}

func hello() string {
	return "hello"
}

func helloMsg(msg string) string {
	return "hello " + msg
}
