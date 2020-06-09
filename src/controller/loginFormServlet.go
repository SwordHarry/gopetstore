package controller

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Hello(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("content-type", "text/html")
	_,err := w.Write([]byte("<html><head></head><body>hello world!</body><html/>"))
	if err != nil {
		panic(err)
	}
}

const loginForm  = "login.html"
const front = "front"

// 返回静态页面
func ShowLogin(w http.ResponseWriter, r *http.Request)  {
	t:=template.Must(template.ParseFiles(filepath.Join(front, loginForm)))

	err := t.Execute(w, nil)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
