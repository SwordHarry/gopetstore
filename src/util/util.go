package util

import (
	"html/template"
	"net/http"
)

// 渲染页面
func Render(w http.ResponseWriter, data interface{}, fileNames ...string) error {
	t := template.Must(template.ParseFiles(fileNames...))
	return t.Execute(w, data)
}
