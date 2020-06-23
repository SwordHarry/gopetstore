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

// file name
const (
	mainFile          = "main.html"
	categoryFile      = "category.html"
	itemFile          = "item.html"
	productFile       = "product.html"
	searchProductFile = "searchProduct.html"
)

// file path
var (
	categoryPath      = filepath.Join(config.Front, config.Web, config.Catalog, categoryFile)
	itemPath          = filepath.Join(config.Front, config.Web, config.Catalog, itemFile)
	mainPath          = filepath.Join(config.Front, config.Web, config.Catalog, mainFile)
	productPath       = filepath.Join(config.Front, config.Web, config.Catalog, productFile)
	searchProductPath = filepath.Join(config.Front, config.Web, config.Catalog, searchProductFile)
)

// about View
// 跳转 主页
func ViewMain(w http.ResponseWriter, r *http.Request) {
	// 跳转到 main主页
	m := make(map[string]interface{})
	err := util.RenderWithAccountAndCommonTem(w, r, m, mainPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 跳转 category 分类页
func ViewCategory(w http.ResponseWriter, r *http.Request) {
	// 获取 get 请求参数
	categoryId := util.GetParam(r, "categoryId")[0]
	c, err := service.GetCategory(categoryId)
	productList, err := service.GetProductList(categoryId)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
	// 从 session 获取 account
	m := make(map[string]interface{})
	m["Category"] = c
	m["ProductList"] = productList
	err = util.RenderWithAccountAndCommonTem(w, r, m, categoryPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 跳转 product 商品页
func ViewProduct(w http.ResponseWriter, r *http.Request) {
	productId := util.GetParam(r, "productId")[0]
	p, err := service.GetProduct(productId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	// 使用 session 存储 product
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("session error: %v", err.Error())
	} else {
		err := s.Save("product", p, w, r)
		if err != nil {
			log.Printf("session error: %v", err.Error())
		}
	}

	items, err := service.GetItemListByProduct(productId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	// render
	m := make(map[string]interface{})
	m["Product"] = p
	m["ItemList"] = items
	err = util.RenderWithAccountAndCommonTem(w, r, m, productPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 跳转 商品 详情页
func ViewItem(w http.ResponseWriter, r *http.Request) {
	// 从 参数中获取 itemid
	itemId := util.GetParam(r, "itemId")[0]
	i, err := service.GetItem(itemId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	// 从 session 中获取 product
	s, err := util.GetSession(r)
	var p *domain.Product
	if err != nil {
		log.Printf("session error: %v", err.Error())
	} else {
		r, ok := s.Get(config.ProductKey)
		if ok {
			p = r.(*domain.Product)
		}
	}
	// render
	m := make(map[string]interface{})
	m["Item"] = i
	m["Product"] = p
	err = util.RenderWithAccountAndCommonTem(w, r, m, itemPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 搜索功能
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("SearchProduct ParseForm error: %v", err.Error())
	}
	keyword := r.FormValue("keyword")
	productList, err := service.SearchProductList("%" + keyword + "%")
	if err != nil {
		log.Printf("SearchProduct SearchProductList error: %v", err.Error())
	}
	m := map[string]interface{}{
		"ProductList": productList,
	}
	err = util.RenderWithAccountAndCommonTem(w, r, m, searchProductPath)
	if err != nil {
		log.Printf("SearchProduct RenderWithAccountAndCommonTem error: %v", err.Error())
	}
}
