package controller

import (
	"gopetstore/src/config"
	"gopetstore/src/util"
	"log"
	"net/http"
	"path/filepath"
)

const mainFile = "main.html"

var mainPath = filepath.Join(config.Front, config.Web, config.Catalog, mainFile)

// 只用于跳转
func ViewMain(w http.ResponseWriter, r *http.Request) {
	// TODO: 从 session 中取出 account
	// 跳转到 main主页
	err := util.RenderWithCommon(w, nil, mainPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
