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

const productFile = "product.html"

var productPath = filepath.Join(config.Front, config.Web, config.Catalog, productFile)

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

	err = util.RenderWithCommon(w, struct {
		Account  interface{}
		Product  *domain.Product
		ItemList []domain.Item
	}{
		Account:  nil,
		Product:  p,
		ItemList: items,
	}, productPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
