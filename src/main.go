package main

import (
	"gopetstore/src/controller"
	"net/http"
)

func main() {
	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("./font"))))
	http.HandleFunc("/showlogin", controller.ShowLogin)
	http.HandleFunc("/login", controller.LoginServlet)
	http.HandleFunc("/hello", controller.Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
