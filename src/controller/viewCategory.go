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

var categoryFile = "category.html"

var categoryPath = filepath.Join(config.Front, config.Web, config.Catalog, categoryFile)

func ViewCategory(w http.ResponseWriter, r *http.Request) {
	// 获取 get 请求参数
	categoryId := util.GetParam(r, "categoryId")[0]
	c, err := service.GetCategory(categoryId)
	productList, err := service.GetProductList(categoryId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	// TODO: Account 待用session 进行存储
	//s,err := util.GetSession(r)
	//if err != nil {
	//	log.Printf("session error: %v",err.Error())
	//}
	err = util.RenderWithCommon(w, struct {
		Account     interface{}
		Category    *domain.Category
		ProductList []domain.Product
	}{
		Account:     nil,
		Category:    c,
		ProductList: productList,
	}, categoryPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
