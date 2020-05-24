package control

import (
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"tools.app/apps/webmvc/model"
)

func TestRegisterIndexPage(t *testing.T) {
	tmpl := template.New("index")
	RegisterFuncMap()
	tmpl.Funcs(GetFuncMap())

	pattern := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/view/*")
	tmpl, err := tmpl.ParseGlob(pattern)
	if err != nil {
		t.Fatal(err)
	}

	user := &model.User{
		ID:       111,
		NickName: "tester01",
		Role:     model.ADMIN,
	}
	err = tmpl.ExecuteTemplate(os.Stdout, tmpl.Name(), user)
	if err != nil {
		t.Fatal(err)
	}
}
