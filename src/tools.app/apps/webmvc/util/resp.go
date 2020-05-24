package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespBody 通用结构体用于装载返回的数据
type RespBody struct {
	Code  int         `json:"code"`
	Rows  interface{} `json:"rows,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

// RespJSON 返回Json的底层方法
func RespJSON(w http.ResponseWriter, data interface{}) {
	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	ret, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	w.Write(ret)
}

// RespOK 操作成功返回Ok
func RespOK(w http.ResponseWriter, data interface{}) {
	RespJSON(w, &RespBody{
		Code: http.StatusOK,
		Data: data,
	})
}

// RespFail 操作失败返回Error
func RespFail(w http.ResponseWriter, failCode int, msg string) {
	RespJSON(w, &RespBody{
		Code: failCode,
		Msg:  msg,
	})
}
