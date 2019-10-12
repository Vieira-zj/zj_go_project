package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golib/httprouter"
	"mock.server/common"
	myutils "tools.app/utils"
)

// MockAPIRegisterHandler register random uri with params and template body.
// Post /mock/register/:uri
func MockAPIRegisterHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// TODO: use db instead of text files.
	uri := params.ByName("uri")
	filePath := fmt.Sprintf("%s/%s_query.txt", common.DataDirPath, uri)
	if err := myutils.WriteContentToFile(filePath, r.URL.RawQuery, true); err != nil {
		common.ErrHandler(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	defer r.Body.Close()

	filePath = fmt.Sprintf("%s/%s_body.txt", common.DataDirPath, uri)
	if err := myutils.WriteContentToFile(filePath, string(body), true); err != nil {
		common.ErrHandler(w, err)
		return
	}

	respJSON := CmdRespJSON{
		Status:  http.StatusOK,
		Message: "success",
		Results: fmt.Sprintf("register uri(%s) success!", uri),
	}
	common.WriteOKJSONResp(w, respJSON)
}

// MockAPIHandler sends templated json response by register params and body.
// Post /mock/:uri
func MockAPIHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uri := params.ByName("uri")
	filePath := fmt.Sprintf("%s/%s_body.txt", common.DataDirPath, uri)
	body, err := myutils.ReadFileContentBuf(filePath)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	defer r.Body.Close()

	tmpl, err := template.New("mockapi").Parse(string(body))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	filePath = fmt.Sprintf("%s/%s_query.txt", common.DataDirPath, uri)
	query, err := myutils.ReadFileContent(filePath)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	tmplParams, err := common.ParseParamsForTempl(common.QueryToMap(query))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, tmplParams); err != nil {
		log.Fatalln(err)
	}
}
