package util

import (
	"gopetstore/src/config"
	"html/template"
	"net/http"
	"path/filepath"
)

// about render
// 渲染页面
func Render(w http.ResponseWriter, data interface{}, fileNames ...string) error {
	_, f := filepath.Split(fileNames[0])
	// 这里传入的 New 中的文件名需要和模板的文件名一致
	// 链式调用，注册 html 片段解析函数
	t, err := template.New(f).
		Funcs(template.FuncMap{"unEscape": UnEscape}).
		ParseFiles(fileNames...)
	if t != nil {
		return t.Execute(w, data)
	}
	return err
}

// 用common模板渲染页面
func RenderWithCommon(w http.ResponseWriter, data interface{}, fileName string) error {
	return Render(w, data, fileName, config.CommonPath)
}

// 将html片段完整输出并要求解析
func UnEscape(s string) template.HTML {
	return template.HTML(s)
}

// 几乎所有页面都用到了 account 信息，故这里再进行 render 的封装
func RenderWithAccount(w http.ResponseWriter, r *http.Request, m map[string]interface{}, fileNames ...string) error {
	// 从 session 获取 account
	a := GetAccountFromSession(r)
	m["Account"] = a
	return Render(w, m, fileNames...)
}

// render with get session from account and render common template
func RenderWithAccountAndCommonTem(w http.ResponseWriter, r *http.Request, m map[string]interface{}, fileName string) error {
	return RenderWithAccount(w, r, m, fileName, config.CommonPath)
}
