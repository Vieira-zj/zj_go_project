package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"src/mock.server/common"
	myutils "src/tools.app/utils"

	"github.com/golib/httprouter"
)

// ToolsHandler router for tools handlers.
func ToolsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	ch := make(chan struct{})

	if r.Method == "POST" {
		switch name {
		case "cmd":
			go runSystemCmd(w, r, ch)
		case "mail":
			go sendMail(w, r, ch)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
	}

	for {
		select {
		case <-ch:
			return
		case <-time.Tick(time.Second):
			log.Println("tools handler is processing...")
		case <-time.After(time.Duration(15) * time.Second):
			common.ErrHandler(w, fmt.Errorf("Time out"))
			return
		}
		time.Sleep(time.Duration(300) * time.Millisecond)
	}
}

// CmdReqJSON run command request json.
type CmdReqJSON struct {
	Meta     string   `json:"meta,omitempty"`
	Commands []string `json:"commands"`
}

// CmdRespJSON run command response json.
type CmdRespJSON struct {
	Meta    string `json:"meta,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Results string `json:"results"`
}

// runSystemCmd, runs shell command and returns results.
// Post /tools/cmd
func runSystemCmd(w http.ResponseWriter, r *http.Request, ch chan<- struct{}) {
	defer func() {
		ch <- struct{}{}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	defer r.Body.Close()

	var reqJSON CmdReqJSON
	if err := json.Unmarshal(body, &reqJSON); err != nil {
		common.ErrHandler(w, err)
		return
	}

	if len(reqJSON.Meta) > 0 {
		log.Println("Run cmd meta:", reqJSON.Meta)
	}
	log.Println("Run cmd:", strings.Join(reqJSON.Commands, ","))

	var result string
	if len(reqJSON.Commands) == 1 {
		if result, err = myutils.RunShellCmdBuf(reqJSON.Commands[0]); err != nil {
			common.ErrHandler(w, err)
			return
		}
	} else {
		if result, err = myutils.RunShellCmds(reqJSON.Commands); err != nil {
			common.ErrHandler(w, err)
			return
		}
	}

	retJSON := CmdRespJSON{
		Status:  http.StatusOK,
		Message: "success",
		Results: result,
	}
	if err := common.WriteOKJSONResp(w, &retJSON); err != nil {
		common.ErrHandler(w, err)
	}
}

// MainRespJSON send mail response json.
type MainRespJSON struct {
	Meta    string `json:"meta,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// sendMail, sends mail.
// Post /tools/mail
func sendMail(w http.ResponseWriter, r *http.Request, ch chan<- struct{}) {
	defer func() {
		ch <- struct{}{}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	defer r.Body.Close()

	entry := myutils.MailEntry{ServerPwd: "*******"}
	if err := json.Unmarshal(body, &entry); err != nil {
		common.ErrHandler(w, err)
		return
	}
	if err := myutils.SendMail(&entry); err != nil {
		common.ErrHandler(w, err)
		return
	}

	retJSON := MainRespJSON{
		Status:  http.StatusOK,
		Message: "success",
	}
	if err := common.WriteOKJSONResp(w, &retJSON); err != nil {
		common.ErrHandler(w, err)
	}
}
