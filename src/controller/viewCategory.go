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
	q := r.URL.Query()
	categoryId := q["categoryId"][0]
	category, err := service.GetCategory(categoryId)
	productList, err := service.GetProductList(categoryId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		// 重定向到 main
		ViewMain(w, r)
		return
	}
	// TODO: Account 待用session 进行存储
	err = util.Render(w, struct {
		Account     interface{}
		Category    domain.Category
		ProductList []domain.Product
	}{
		Account:     nil,
		Category:    *category,
		ProductList: productList,
	}, categoryPath, config.Common)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
