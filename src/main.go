package main

import (
	"gopetstore/src/route"
	"net/http"
)

const port = ":8080"

func main() {
	// 静态文件服务器
	http.Handle("/", http.FileServer(http.Dir("front")))
	// 注册路由
	route.RegisterRoute()
	// 监听端口
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
