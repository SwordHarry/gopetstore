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
	q := r.URL.Query()
	productId := q["productId"][0]
	p, err := service.GetProduct(productId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		ViewCategory(w, r)
		return
	}
	items, err := service.GetItemListByProduct(productId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		ViewCategory(w, r)
		return
	}

	err = util.Render(w, struct {
		Account  interface{}
		Product  domain.Product
		ItemList []domain.Item
	}{
		Account:  nil,
		Product:  *p,
		ItemList: items,
	}, productPath, config.Common)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
