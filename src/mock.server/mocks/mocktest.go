package mocks

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	redis "gopkg.in/redis.v5"
)

// client test scripts in ex_http.go

// MockTest1 : parse get request, => /test1?userid="xxx"&username="xxx"&url=url1&url=url2
func MockTest1(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("\nREQUEST DUMP ERROR:", err.Error())
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

// MockTest2 : parse post request, => /test2 + json_body
func MockTest2(w http.ResponseWriter, r *http.Request) {
	printRequestData(r)

	result, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
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

// MockTest3 : get post request data, => /test3 + header + body(k=v)
func MockTest3(w http.ResponseWriter, r *http.Request) {
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

// MockTest4 : test get access count from redis
func MockTest4(rw http.ResponseWriter, req *http.Request) {
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	log.Println("***** Configs *****\nredis server:", RunConfigs.Server.Redis)
	accessCount, err := getAccessCount()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("TOTAL ACCESS:", accessCount)

	retContent := fmt.Sprintln("Access count: " + accessCount)
	rw.Header().Set("Content-Length", strconv.Itoa(len([]byte(retContent))))
	rw.WriteHeader(http.StatusOK)
	log.Println("return 200")

	io.Copy(rw, strings.NewReader(retContent))
	log.Print("===> MockTest4, send data done\n\n")
}

func getAccessCount() (string, error) {
	options := redis.Options{
		Addr:     RunConfigs.Server.Redis,
		Password: "",
	}
	client := redis.NewClient(&options)
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		return "-1", err
	}
	log.Println(pong)

	const key = "mock_total_access"
	total, err := client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			if err := client.Set(key, 1, 0).Err(); err != nil {
				return "-1", err
			}
			total = "1"
		} else {
			return "-1", err
		}
	}

	err = client.Incr(key).Err()
	if err != nil {
		return "-1", err
	}

	return total, nil
}
