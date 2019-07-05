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
	"sync"
	"time"
)

// demo 01, http get request
func testHTTPGet() {
	u, _ := url.Parse("http://127.0.0.1:17891/test/1")
	q := u.Query()
	q.Set("userid", "idxxx")
	q.Set("username", "namexxx")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("\nhttp get resp body:\n", string(result))
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

func testHTTPPost() {
	var s serverList
	s.ServersID = "level01"
	s.Servers = append(s.Servers, server{ServerName: "GuangZhou", ServerIP: "127.0.0.10"})
	s.Servers = append(s.Servers, server{ServerName: "ShangHai", ServerIP: "127.0.0.11"})
	s.Servers = append(s.Servers, server{ServerName: "BeiJing", ServerIP: "127.0.0.13"})

	b, err := json.Marshal(&s)
	if err != nil {
		panic(err.Error())
	}
	body := bytes.NewBuffer(b)

	const url = "http://127.0.0.1:17891/test/2"
	resp, err := http.Post(url, "application/json;sharset=uft-8", body)
	if err != nil {
		panic(err.Error())
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("\nhttp post resp body:\n", string(result))
}

// demo 03, post with http client
func testClientHTTP() {
	const url = "http://127.0.0.1:17891/test/3"
	req, err := http.NewRequest("POST", url, strings.NewReader("key=val"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "test_c=test_val")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nhttp post resp body:\n", string(body))
}

// demo 04, read http stream data
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

func testHTTPStreamRead() {
	const threads = 10
	var wg sync.WaitGroup
	args := consumeMsgArgs{
		User:      "user",
		QueueName: "queueName",
		Stream:    true,
		Limit:     "100",
		// Position:  "",
	}

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, idx int) {
			defer wg.Done()
			httpStreamRead(args, "zj_test_verified_msg")
			fmt.Printf("***** stream consume thread-%d done\n", idx)
		}(&wg, i)
	}

	wg.Wait()
	fmt.Println("read http stream data done.")
}

func httpStreamRead(args consumeMsgArgs, verifiedMsg string) {
	u := "http://10.200.20.36:14532" + "/queues/" + args.QueueName + "/consume"
	v := url.Values{}
	v.Add("stream", strconv.FormatBool(args.Stream))
	if len(args.Limit) > 0 {
		v.Add("limit", args.Limit)
	}
	if len(args.Position) > 0 {
		v.Add("position", args.Position)
	}
	url := u + "?" + v.Encode()

	// fmt.Println("***** request url:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "QBox token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("***** Status:", resp.StatusCode)
	fmt.Printf("***** Response header %v\n", resp.Header)
	if resp.StatusCode != 200 {
		time.Sleep(time.Second)
		output, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		if len(output) > 0 {
			fmt.Println("***** Error message:", string(output))
		}
		return
	}

	// stream read
	// one client produce stream data, and another client consume data sync
	ch := make(chan bool)
	go func() {
		fmt.Println("***** Stream data:")
		total := 0
		decStream := json.NewDecoder(resp.Body)
		for {
			total++
			var msg streamMsg
			err := decStream.Decode(&msg)
			if err != nil {
				fmt.Println(err.Error())
				ch <- false
				return
			}
			if msg.Message == verifiedMsg {
				ch <- true
				return
			}
			if total%1000 == 0 {
				fmt.Printf("***** iterator at %d\n", total)
				fmt.Printf("{position: %s, msg: %s}\n", msg.Position, msg.Message)
			}
			// time.Sleep(800 * time.Millisecond)
		}
	}()

	// stream read connection will not be closed by server,
	// and here set a timeout(60s) to close connection.
	select {
	case <-time.After(60 * time.Second):
		fmt.Println("***** stream read timeout.")
	case ret := <-ch:
		if ret {
			fmt.Println("***** message found.")
		} else {
			fmt.Println("***** message not found.")
		}
	}
}

// MainHTTP : main function, http client testing for mock server.
func MainHTTP() {
	// testHTTPGet()
	// testHTTPPost()
	// testClientHTTP()

	// testHTTPStreamRead()

	fmt.Println("golang http client example DONE.")
}
