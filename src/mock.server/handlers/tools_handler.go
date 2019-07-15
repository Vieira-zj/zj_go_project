package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golib/httprouter"
	"mock.server/common"
	myutils "tools.app/utils"
)

// ToolsHandler router for tools handlers.
func ToolsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	if r.Method == "POST" {
		switch name {
		case "cmd":
			runSystemCmd(w, r)
		case "mail":
			sendMail(w, r)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
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
func runSystemCmd(w http.ResponseWriter, r *http.Request) {
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

// sendMail, sends mail.
// Post /tools/mail
func sendMail(w http.ResponseWriter, r *http.Request) {
}
