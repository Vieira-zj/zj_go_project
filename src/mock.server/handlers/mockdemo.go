package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golib/httprouter"
	redis "gopkg.in/redis.v5"
	"mock.server/common"
)

// MockDemoHandler router for mock demos.
func MockDemoHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer func() {
		if p := recover(); p != nil {
			common.ErrHandler(w, p.(error))
		}
	}()

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockDemo01(w, r, params)
		case 2:
			mockDemo02(w, r, params)
		case 5:
			mockDemo05(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
			return
		}
	}

	if r.Method == "POST" {
		switch id {
		case 3:
			mockDemo03(w, r, params)
		case 4:
			mockDemo04(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("POST for invalid path: %s", r.URL.Path))
		}
	}
}

// demo, parse get request => Get /demo/01?userid=xxx&username=xxx
func mockDemo01(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}

	r.ParseForm()
	log.Println("Request Method:", r.Method)
	log.Println("Form Data:")
	log.Println("userid:", r.Form["userid"][0])
	log.Println("username:", r.Form["username"][0])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

// demo, parse get request => Get /demo/02?userid=xxx&username=xxx&key=val1&key=val2
func mockDemo02(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}

	values := r.URL.Query()
	log.Println("Request Query:")
	fmt.Println("userid:", values["userid"][0])
	fmt.Println("username:", values["username"][0])
	for _, v := range values["key"] {
		fmt.Println("key:", v)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

// demo, parse post json body
type server struct {
	ServerName string `json:"server_name"`
	ServerIP   string `json:"server_ip"`
}

type serverInfo struct {
	SvrList  []server `json:"server_list"`
	SvrGrpID string   `json:"server_group_id"`
}

// => Post /demo/03
func mockDemo03(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	printServerInfo(body)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

func printServerInfo(jsonBody []byte) {
	var s serverInfo
	json.Unmarshal(jsonBody, &s)
	log.Printf("Request body: %+v\n", s)

	log.Println("Server Info:")
	log.Println("server group id:", s.SvrGrpID)
	for _, svr := range s.SvrList {
		log.Printf("server name: %s, server ip: %s\n", svr.ServerName, svr.ServerIP)
	}
}

// demo, parse post form with cookie => POST /demo/04
func mockDemo04(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}

	log.Println("Request Data:")
	log.Printf("content type: %s\n", r.Header.Get(common.TextContentType))
	log.Printf("form kv: key1=%s\n", r.PostFormValue("key1"))
	log.Printf("form kv: key2=%s\n", r.PostFormValue("key2"))

	cookie, err := r.Cookie("user")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	log.Println("name:", cookie.Name)
	log.Println("value:", cookie.Value)

	cookie, err = r.Cookie("pwd")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	log.Println("name:", cookie.Name)
	log.Println("value:", cookie.Value)
	log.Println("domain:", cookie.Domain)
	log.Println("expires:", cookie.Expires)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hi, thanks for access %s", html.EscapeString(r.URL.Path[1:]))
}

// demo, test get access count by redis => GET /demo/05
// redis env: docker run --name redis -p 6379:6379 --rm -d redis:4.0
func mockDemo05(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}

	if len(common.RunConfigs.Server.RedisURI) == 0 {
		common.ErrHandler(w, fmt.Errorf("config redis uri is empty"))
		return
	}

	log.Println("redis server:", common.RunConfigs.Server.RedisURI)
	accessCount, err := getAccessCount()
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	log.Println("Total Access:", accessCount)

	retContent := fmt.Sprintln("Total Access: " + accessCount)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len([]byte(retContent))))
	w.WriteHeader(http.StatusOK)
	io.Copy(w, strings.NewReader(retContent))
}

func getAccessCount() (string, error) {
	errNum := "-1"
	options := redis.Options{
		Addr:     common.RunConfigs.Server.RedisURI,
		Password: "",
	}

	client := redis.NewClient(&options)
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		return errNum, err
	}
	log.Println(pong)

	const key = "server_total_access"
	total, err := client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			if err := client.Set(key, 1, 0).Err(); err != nil {
				return errNum, err
			}
			total = "1"
		} else {
			return errNum, err
		}
	}

	err = client.Incr(key).Err()
	if err != nil {
		return "-1", err
	}

	return total, nil
}
