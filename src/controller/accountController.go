package controller

import (
	"gopetstore/src/config"
	"gopetstore/src/domain"
	"gopetstore/src/service"
	"gopetstore/src/util"
	"log"
	"net/http"
	"path/filepath"
)

const signInFormFile = "signInForm.html"
const registerFormFile = "registerForm.html"
const accountFieldFile = "accountField.html"

var signInFormPath = filepath.Join(config.Front, config.Web, config.Account, signInFormFile)
var registerFormPath = filepath.Join(config.Front, config.Web, config.Account, registerFormFile)
var accountFieldPath = filepath.Join(config.Front, config.Web, config.Account, accountFieldFile)

// 跳转到登录页面
func ViewLoginOrPostLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := util.RenderWithCommon(w, nil, signInFormPath)
		if err != nil {
			log.Printf("view signInForm error: %v", err.Error())
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Printf("parse login form error: %v", err.Error())
			return
		}
		userName := r.FormValue("username")
		password := r.FormValue("password")
		account, err := service.GetAccountByUserNameAndPassword(userName, password)
		if err != nil {
			log.Printf("do login error with %s %s: %v", userName, password, err.Error())
			return
		}
		if account != nil {
			// session 中保存 account
			s, err := util.GetSession(r)
			if err != nil {
				log.Printf("get session error: %v", err.Error())
				return
			}
			err = s.Save("account", account, w, r)
			if err != nil {
				log.Printf("session save account error: %v", err.Error())
				return
			}
			// 登录成功后跳转到主页
			ViewMain(w, r)
		}
	}
}

// 跳转到注册页面
func ViewRegister(w http.ResponseWriter, r *http.Request) {
	languages := []string{
		"english",
		"japanese",
	}
	categories := []string{
		"FISH",
		"DOGS",
		"REPTILES",
		"CATS",
		"BIRDS",
	}
	err := util.Render(w, struct {
		Account    *domain.Account
		Languages  []string
		Categories []string
	}{
		Account:    nil,
		Languages:  languages,
		Categories: categories,
	}, registerFormPath, config.CommonPath, accountFieldPath)
	if err != nil {
		log.Printf("view registerForm error: %v", err.Error())
	}
}
