package httptemplate

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	myutil "tools.app/utils"
)

// TemplateHandler01 returns hello world template html.
func TemplateHandler01(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/01_layout.html")
	if err != nil {
		myutil.WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("[TemplateHandler01]:", err)
		return
	}

	fmt.Println("use template:", tmpl.Name())
	if err := tmpl.Execute(w, "Hello world"); err != nil {
		myutil.WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("[TemplateHandler01]:", err)
	}
}

// TemplateHandler02 returns layout.html and index.html template html.
func TemplateHandler02(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/02_layout.html", "templates/02_index.html")
	if err != nil {
		myutil.WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("[TemplateHandler01]:", err)
		return
	}

	const tmplName = "02_layout"
	fmt.Println("use template:", tmplName)
	if err := t.ExecuteTemplate(w, tmplName, "Hello world"); err != nil {
		myutil.WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("[TemplateHandler01]:", err)
	}
}
