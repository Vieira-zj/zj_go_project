package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// demo 01, http get request
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

// demo 02, http post request
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

	const url = "http://localhost:17890/mock02"
	resp, err := http.Post(url, "application/json;sharset=uft-8", body)
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

// demo 03, post with http client
func testClientHTTP() {
	const url = "http://localhost:17890/mock03"
	req, err := http.NewRequest("POST", url, strings.NewReader("key=test"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "test_k=test_c")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("response body:\n%s\n", string(body))
}

// demo 04, read stream data function
type consumeMsgArgs struct {
	User      string
	QueueName string
	Position  string
	Limit     string
	Stream    bool
}

type streamMsg struct {
	Position string `json:"position"`
	Message  string `json:"message"`
}

func streamRead(args consumeMsgArgs, limit int, verifiedMsg string) {
	u := "http://10.200.20.38:12500/kmq/queues/" + args.QueueName + "/consume"
	v := url.Values{}
	if len(args.Position) > 0 {
		v.Add("position", args.Position)
	}
	v.Add("stream", strconv.FormatBool(true))
	url := u + "?" + v.Encode()

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "QBox "+args.User)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("***** Response status code:", resp.StatusCode)
	if resp.StatusCode != 200 {
		output, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("***** Response err message:", string(output))
		return
	}

	// stream read
	decStream := json.NewDecoder(resp.Body)
	fmt.Println("***** Stream data:")
	for i := 0; i < limit; i++ {
		var msg streamMsg
		err := decStream.Decode(&msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("{position: %s, msg: %s}\n", msg.Position, msg.Message)
		if msg.Message == verifiedMsg {
			fmt.Println("***** Message found!")
			return
		}
		// time.Sleep(300 * time.Millisecond)
	}
}

// MainHTTP : main for the http client test demos for mock server.
func MainHTTP() {
	// testClientGet()
	// testClientPost()
	testClientHTTP()

	fmt.Println("http client demo.")
}
