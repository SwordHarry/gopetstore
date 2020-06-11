package util

import (
	"gopetstore/src/config"
	"html/template"
	"net/http"
)

// about render
// 渲染页面
func Render(w http.ResponseWriter, data interface{}, fileNames ...string) error {
	t := template.Must(template.ParseFiles(fileNames...))
	return t.Execute(w, data)
}

// 用common模板渲染页面
func RenderWithCommon(w http.ResponseWriter, data interface{}, fileName string) error {
	return Render(w, data, fileName, config.CommonPath)
}

// 将html片段完整输出并要求解析
func UnEscape(s string) template.HTML {
	return template.HTML(s)
}
