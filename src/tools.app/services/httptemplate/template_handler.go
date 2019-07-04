package httptemplate

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	myutil "tools.app/utils"
)

// TemplateHandler router for template handlers.
func TemplateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idx, err := strconv.Atoi(params.ByName("idx"))
	if err != nil {
		errHandler(w, err)
		return
	}

	// 在switch分支语句的每个case中会自动加上一个break语句
	switch idx {
	case 1:
		TemplateHandler01(w, r)
	case 2:
		TemplateHandler02(w, r)
	case 3:
		TemplateHandler03(w, r)
	case 4:
		TemplateHandler04(w, r)
	default:
		errHandler(w, fmt.Errorf("path template/%d not exist", idx))
	}
}

// TemplateHandler01 test hello world html.
func TemplateHandler01(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/01_layout.html")
	if err != nil {
		errHandler(w, err)
		return
	}

	fmt.Println("use template:", tmpl.Name())
	if err := tmpl.Execute(w, "Hello world"); err != nil {
		errHandler(w, err)
	}
}

// TemplateHandler02 模板命名与嵌套
func TemplateHandler02(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/02_layout.html", "templates/02_index.html")
	if err != nil {
		errHandler(w, err)
		return
	}

	const tmplName = "02_layout"
	fmt.Println("use template:", tmplName)
	if err := t.ExecuteTemplate(w, tmplName, "Hello world"); err != nil {
		errHandler(w, err)
	}
}

// TemplateHandler03 test template statements "range" and "with".
func TemplateHandler03(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/03_layout.html")
	if err != nil {
		errHandler(w, err)
		return
	}

	const tmplName = "03_layout"
	fmt.Println("use template:", tmplName)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Ths", "Fri", "Sat", "Sun"}
	if err := tmpl.ExecuteTemplate(w, tmplName, daysOfWeek); err != nil {
		errHandler(w, err)
	}
}

// TemplateHandler04 智能上下文
func TemplateHandler04(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/04_layout.html")
	if err != nil {
		errHandler(w, err)
		return
	}

	const tmplName = "04_layout"
	fmt.Println("use template:", tmplName)
	content := `I asked: <i>What's up?</i>`
	if err := tmpl.ExecuteTemplate(w, tmplName, content); err != nil {
		errHandler(w, err)
	}
}

func errHandler(w http.ResponseWriter, err error) {
	myutil.WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
	fmt.Println(err)
}
