package controller

import (
	"gopetstore/src/config"
	"gopetstore/src/domain"
	"gopetstore/src/service"
	"gopetstore/src/util"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const itemFile = "item.html"

var itemPath = filepath.Join(config.Front, config.Web, config.Catalog, itemFile)

func ViewItem(w http.ResponseWriter, r *http.Request) {
	itemId := util.GetParam(r, "itemId")[0]
	i, err := service.GetItem(itemId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	s, err := util.GetSession(r)
	var p *domain.Product
	if err != nil {
		log.Printf("session error: %v", err.Error())
	} else {
		r, ok := s.Get("product")
		if ok {
			p = r.(*domain.Product)
		}
	}

	var descriptionHtml template.HTML
	if p != nil {
		descriptionHtml = util.UnEscape(p.Description)
	}
	err = util.RenderWithCommon(w, struct {
		Account  interface{}
		Item     *domain.Item
		Product  *domain.Product
		DescHTML template.HTML
	}{
		Account:  nil,
		Item:     i,
		Product:  p,
		DescHTML: descriptionHtml,
	}, itemPath)

	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
