package mocks

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// client test scripts in ex_http.go

// Mock21 : parse get request, => /mock1?userid="xxx"&username="xxx"&url=url1&url=url2
func Mock21(w http.ResponseWriter, r *http.Request) {
	printRequestData(r)

	// #1
	// r.ParseForm()
	// fmt.Println("method:", r.Method)
	// fmt.Println("userid:", r.Form["userid"][0])
	// fmt.Println("username:", r.Form["username"][0])

	// #2
	values := r.URL.Query()
	fmt.Println("userid:", values["userid"][0])
	fmt.Println("username:", values["username"][0])
	if len(values["url"]) > 0 {
		for _, v := range values["url"] {
			fmt.Println("url:", v)
		}
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

func printRequestData(r *http.Request) {
	reqHeader, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nREQUEST:\n%s\n", reqHeader)
}

// struct for post json body
type server struct {
	ServerName string `json:"servers_name"`
	ServerIP   string `json:"servers_ip"`
}

type serverList struct {
	Servers   []server `json:"servers_list"`
	ServersID string   `json:"servers_group_id"`
}

// Mock22 : parse post request, => /mock2 + json_body
func Mock22(w http.ResponseWriter, r *http.Request) {
	printRequestData(r)

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	printServersInfoByStruct(result)
	printServersInfoByInterface(result)

	w.WriteHeader(200)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

func printServersInfoByStruct(jsonBody []byte) {
	var s serverList
	json.Unmarshal(jsonBody, &s)
	fmt.Println("REQUEST JSON BODY:")
	fmt.Println("servers group id:", s.ServersID)
	for _, ser := range s.Servers {
		fmt.Printf("server name: %s, server ip: %s\n", ser.ServerName, ser.ServerIP)
	}
}

func printServersInfoByInterface(jsonBody []byte) {
	var f interface{}
	json.Unmarshal(jsonBody, &f)
	m := f.(map[string]interface{})
	fmt.Println("REQUEST JSON BODY:")
	fmt.Println("servers group id:", m["servers_group_id"])

	f = m["servers_list"]
	for _, s := range f.([]interface{}) {
		fmt.Printf("server name: %s, server ip: %s\n",
			s.(map[string]interface{})["servers_name"],
			s.(map[string]interface{})["servers_ip"])
	}
}

// Mock23 : get post request data, => /mock3 + header + body(k=v)
func Mock23(w http.ResponseWriter, r *http.Request) {
	printRequestData(r)

	fmt.Printf("\ncontent type: %s\n", r.Header.Get("Content-Type"))
	fmt.Printf("post kv: key=%s\n", r.PostFormValue("key"))

	w.WriteHeader(200)
	cookie, err := r.Cookie("test_k")
	if err == nil {
		fmt.Fprintln(w, "domain:", cookie.Domain)
		fmt.Fprintln(w, "expires:", cookie.Expires)
		fmt.Fprintln(w, "name:", cookie.Name)
		fmt.Fprintln(w, "value:", cookie.Value)
	}
}
