package controller

import (
	"gopetstore/src/domain"
	"gopetstore/src/service"
	"net/http"
)


func LoginServlet(w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	u := domain.User{
		UserName: username,
		Password: password,
	}
	up := service.Login(&u)
	if up != nil {
		_,_ = w.Write([]byte("<html><head></head><body>success</body><html/>"))
	} else {
		_,_ = w.Write([]byte("<html><head></head><body>failed</body><html/>"))
	}
}
