package control

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"src/tools.app/apps/webmvc/model"
	"src/tools.app/apps/webmvc/util"
)

// RegisterView 注册视图
func RegisterView() {
	RegisterFuncMap()
	RegisterIndexView()
	// RegisterXxxView()
}

// RegisterIndexView 注册Index视图
func RegisterIndexView() {
	tplname := "index"
	tmpl := template.New(tplname)
	tmpl.Funcs(GetFuncMap())

	glob := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/view/*")
	tmpl, err := tmpl.ParseGlob(glob)
	if err != nil {
		log.Fatal(err)
	}

	pattern := "/" + tplname
	log.Printf("register template view %s:%s\n", pattern, tplname)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.ExecuteTemplate(w, tplname, loadUserFromSession(req))
		if err != nil {
			util.RespFail(w, http.StatusInternalServerError, err.Error())
		}
	})
	return
}

func loadUserFromSession(req *http.Request) *model.User {
	// mock
	return &model.User{
		ID:       111,
		NickName: "mockUser01",
		Role:     model.ADMIN,
	}
}

// RegisterCtrl 注册控制器
func RegisterCtrl() {
	RegisterRegExRouter()
	new(UserCtrl).Router()
	// new(OpenCtrl).Router()
	// new(AttachCtrl).Router()
}

// RouterGet 添加get路由规则
func RouterGet(pattern string, fun func(w http.ResponseWriter, req *http.Request)) {
	log.Println("register router (get):", pattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			fun(w, req)
		} else {
			util.RespFail(w, http.StatusInternalServerError,
				fmt.Sprintf("not support uri=%s, method=%s\n", req.URL, req.Method))
		}
	})
}

// RouterPost 添加post路由规则
func RouterPost(pattern string, fun func(w http.ResponseWriter, req *http.Request)) {
	log.Println("register router (post):", pattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			fun(w, req)
		} else {
			util.RespFail(w, http.StatusInternalServerError,
				fmt.Sprintf("not support uri=%s, method=%s\n", req.URL, req.Method))
		}
	})
}

// RegExRouterMap 保存uri和处理函数之间关系的字典
var RegExRouterMap = make(map[string]func(w http.ResponseWriter, req *http.Request))

// RegexpMatchMap 保存uri和对应正则表达式的字典
var RegexpMatchMap = make(map[string]*regexp.Regexp)

// RegExRouter 添加路由规则, 支持正则
func RegExRouter(pattern string, fun func(w http.ResponseWriter, req *http.Request)) {
	RegExRouterMap[pattern] = fun
	RegexpMatchMap[pattern], _ = regexp.Compile(pattern)
}

// RegisterRegExRouter 全局通过正则将url和处理函数绑定
func RegisterRegExRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		uris := strings.Split(req.RequestURI, "?")
		uri := uris[0]
		handlefunc := notfound
		for p, match := range RegexpMatchMap {
			if match.MatchString(uri) {
				handlefunc = RegExRouterMap[p]
				break
			}
		}
		handlefunc(w, req)
	})
}

func notfound(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("RESOURCES NOT FOUND"))
}
