package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Bind 请求参数绑定
func Bind(req *http.Request, obj interface{}) error {
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		return BindJSON(req, obj)
	}
	if strings.Contains(strings.ToLower(contentType), "application/x-www-form-urlencoded") {
		return BindForm(req, obj)
	}
	if strings.Contains(strings.ToLower(contentType), "text/xml") {
		return BindXML(req, obj)
	}
	return fmt.Errorf("Content-Type %s not support", contentType)
}

// BindJSON json参数绑定
func BindJSON(req *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}

// BindForm from参数绑定
func BindForm(req *http.Request, obj interface{}) error {
	// mock
	return nil
}

// BindXML xml参数绑定
func BindXML(req *http.Request, obj interface{}) error {
	// mock
	return nil
}
