package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// demo 01, get request
func testClientGet() {
	u, _ := url.Parse("http://localhost:17890/mock01")
	q := u.Query()
	q.Set("userid", "tester01")
	q.Set("username", "tester_a")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("response body: %s\n", string(result))
}

// demo 02, post request
type server struct {
	ServerName string `json:"servers_name"`
	ServerIP   string `json:"servers_ip"`
}

type serverList struct {
	Servers   []server `json:"servers_list"`
	ServersID string   `json:"servers_group_id"`
}

func testClientPost() {
	var s serverList
	s.ServersID = "group01"
	s.Servers = append(s.Servers, server{ServerName: "GuangZhou", ServerIP: "127.0.0.10"})
	s.Servers = append(s.Servers, server{ServerName: "ShangHai", ServerIP: "127.0.0.11"})
	s.Servers = append(s.Servers, server{ServerName: "BeiJing", ServerIP: "127.0.0.13"})

	b, err := json.Marshal(&s)
	if err != nil {
		panic(err.Error())
	}
	body := bytes.NewBuffer(b)

	resp, err := http.Post("http://localhost:17890/mock02", "application/json;sharset=uft-8", body)
	if err != nil {
		panic(err.Error())
	}
	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("response body: %s\n", string(result))
}

// demo 03, post with client http
func testClientHTTP() {
	req, err := http.NewRequest("POST", "http://localhost:17890/mock03", strings.NewReader("key=test"))
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "test_k=test_c")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	fmt.Printf("response body:\n%s\n", string(body))
}

// MainHTTP : main for the http client test demos for mock server.
func MainHTTP() {
	// testClientGet()
	// testClientPost()
	testClientHTTP()

	fmt.Println("http client demo.")
}
