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

const (
	signInFormFile      = "signInForm.html"
	registerFormFile    = "registerForm.html"
	accountFieldFile    = "accountField.html"
	editAccountFormFile = "editAccountForm.html"
)

var (
	signInFormPath      = filepath.Join(config.Front, config.Web, config.Account, signInFormFile)
	registerFormPath    = filepath.Join(config.Front, config.Web, config.Account, registerFormFile)
	accountFieldPath    = filepath.Join(config.Front, config.Web, config.Account, accountFieldFile)
	editAccountFormPath = filepath.Join(config.Front, config.Web, config.Account, editAccountFormFile)
	languages           = []string{
		"english",
		"japanese",
	}
	categories = []string{
		"FISH",
		"DOGS",
		"REPTILES",
		"CATS",
		"BIRDS",
	}
)

// get 跳转到登录页面 或者 post 登录
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
		}
		userName := r.FormValue("username")
		password := r.FormValue("password")
		account, err := service.GetAccountByUserNameAndPassword(userName, password)
		if err != nil {
			log.Printf("do login error with %s %s: %v", userName, password, err.Error())
			m := map[string]interface{}{
				"Message": "登录失败,用户名或密码错误！",
			}
			err := util.RenderWithAccountAndCommonTem(w, r, m, signInFormPath)
			if err != nil {
				log.Printf("view signInForm error: %v", err.Error())
			}
		}
		if account != nil {
			// session 中保存 account
			s, err := util.GetSession(r)
			if err != nil {
				log.Printf("get session error: %v", err.Error())
			}
			if s != nil {
				err = s.Save("account", account, w, r)
				if err != nil {
					log.Printf("session save account error: %v", err.Error())
				}
				// 登录成功后跳转到主页
				ViewMain(w, r)
			}
		}
	}
}

// 跳转到注册页面
func ViewRegister(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]interface{})
	m["Languages"] = languages
	m["Categories"] = categories
	err := util.RenderWithAccount(w, r, m, registerFormPath, config.CommonPath, accountFieldPath)
	if err != nil {
		log.Printf("view registerForm error: %v", err.Error())
	}
}

// 登出
func SignOut(w http.ResponseWriter, r *http.Request) {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
	}
	if s != nil {
		err = s.Del("account", w, r)
		err = s.Del("cart", w, r)
		err = s.Del("order", w, r)
		if err != nil {
			log.Printf("session delete error: %v", err.Error())
		}
	}
	// 重定向到主页
	ViewMain(w, r)
}

// 编辑用户信息
func ViewEditAccount(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	m["Languages"] = languages
	m["Categories"] = categories
	err := util.RenderWithAccount(w, r, m, editAccountFormPath, config.CommonPath, accountFieldPath)
	if err != nil {
		log.Printf("view editAccountForm error: %v", err.Error())
	}
}

// 注册用户
func NewAccount(w http.ResponseWriter, r *http.Request) {
	a := getAccountFromInfoForm(r)
	repeatedPassword := r.FormValue("repeatedPassword")
	m := make(map[string]interface{})
	m["Languages"] = languages
	m["Categories"] = categories
	if len(a.Password) == 0 {
		m["Account"] = a
		m["Message"] = "密码不能为空"
		// 返回到注册页面
		err := util.Render(w, m, registerFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("NewAccount RenderWithAccount error: %v", err.Error())
		}
		return
	}
	if repeatedPassword != a.Password {
		m["Account"] = a
		m["Message"] = "密码和重复密码不符"
		// 返回到注册页面
		err := util.Render(w, m, registerFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("NewAccount RenderWithAccount error: %v", err.Error())
		}
		return
	}
	// 检查是否已存在
	oldAccount, err := service.GetAccountByUserName(a.UserName)
	if oldAccount != nil {
		m["Account"] = a
		m["Message"] = "用户名已存在"
		// 返回到注册页面
		err := util.Render(w, m, registerFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("NewAccount RenderWithAccount error: %v", err.Error())
		}
		return
	}
	// 进行注册
	err = service.InsertAccount(a)
	if err != nil {
		log.Printf("NewAccount InsertAccount error: %v", err.Error())
	}
	// 到登录页面并提示注册成功
	m["Message"] = "注册成功!"
	err = util.RenderWithCommon(w, m, signInFormPath)
	if err != nil {
		log.Printf("NewAccount RenderWithAccountAndCommonTem error: %v", err.Error())
	}
}

// 确认修改用户信息
func ConfirmEdit(w http.ResponseWriter, r *http.Request) {
	a := getAccountFromInfoForm(r)
	repeatedPassword := r.FormValue("repeatedPassword")
	m := make(map[string]interface{})
	m["Languages"] = languages
	m["Categories"] = categories
	if len(a.Password) == 0 {
		m["Message"] = "密码不能为空"
		err := util.RenderWithAccount(w, r, m, editAccountFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("ConfirmEdit RenderWithAccount error: %v", err.Error())
		}
		return
	}
	if repeatedPassword != a.Password {
		m["Message"] = "密码和重复密码不符"
		// 返回到修改信息页面
		err := util.RenderWithAccount(w, r, m, editAccountFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("ConfirmEdit RenderWithAccount error: %v", err.Error())
		}
		return
	}
	err := service.UpdateAccount(a)
	if err != nil {
		log.Printf("ConfirmEdit UpdateAccount error: %v", err.Error())
		m["Message"] = err.Error()
		// 返回到修改信息页面
		err := util.RenderWithAccount(w, r, m, editAccountFormPath, config.CommonPath, accountFieldPath)
		if err != nil {
			log.Printf("ConfirmEdit UpdateAccount error: %v", err.Error())
		}
		return
	}
	newAccount, err := service.GetAccountByUserName(a.UserName)
	if err != nil {
		log.Printf("ConfirmEdit GetAccountByUserName error: %v", err.Error())
	}
	m["Message"] = "修改成功"
	// 修改成功后需要重置 session
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ConfirmEdit GetSession error: %v", err.Error())
	}
	if s != nil {
		err = s.Save("account", newAccount, w, r)
		if err != nil {
			log.Printf("ConfirmEdit Save error: %v", err.Error())
		}
	}
	// 返回到修改信息页面
	err = util.RenderWithAccount(w, r, m, editAccountFormPath, config.CommonPath, accountFieldPath)
	if err != nil {
		log.Printf("ConfirmEdit RenderWithAccount error: %v", err.Error())
	}
}

// 从用户详情表单中获取 account
func getAccountFromInfoForm(r *http.Request) *domain.Account {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse register form error: %v", err.Error())
	}
	userName := r.FormValue("username")
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address1 := r.FormValue("address1")
	address2 := r.FormValue("address2")
	city := r.FormValue("city")
	state := r.FormValue("state")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	languagePreference := r.FormValue("languagePreference")
	favouriteCategoryId := r.FormValue("favouriteCategoryId")
	listOption := r.FormValue("listOption")
	bannerOption := r.FormValue("bannerOption")

	finalListOption := len(listOption) > 0
	finalBannerOption := len(bannerOption) > 0
	a := &domain.Account{
		UserName:            userName,
		Password:            password,
		Email:               email,
		FirstName:           firstName,
		LastName:            lastName,
		Address1:            address1,
		Address2:            address2,
		City:                city,
		State:               state,
		Zip:                 zip,
		Country:             country,
		Phone:               phone,
		FavouriteCategoryId: favouriteCategoryId,
		LanguagePreference:  languagePreference,
		ListOption:          finalListOption,
		BannerOption:        finalBannerOption,
	}
	return a
}
