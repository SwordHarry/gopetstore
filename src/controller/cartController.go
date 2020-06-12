package controller

import (
	"gopetstore/src/config"
	"gopetstore/src/domain"
	"gopetstore/src/service"
	"gopetstore/src/util"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

const cartFile = "cart.html"

var cartPath = filepath.Join(config.Front, config.Web, config.Cart, cartFile)

// 跳转到购物车页面
func ViewCart(w http.ResponseWriter, r *http.Request) {
	cart := util.GetCartFromSession(w, r, func(cart *domain.Cart) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Printf("parse form error: %v", err.Error())
				return
			}
			for _, ci := range cart.ItemList {
				itemId := ci.Item.ItemId
				quantityStr := r.PostFormValue(itemId)
				quantity, err := strconv.Atoi(quantityStr)
				if err != nil {
					log.Printf("set quantity error: %v", err.Error())
					continue
				}
				log.Printf("set quantity: %v, %v", itemId, quantity)
				cart.SetQuantityByItemId(itemId, quantity)
			}
		}
	})

	// render
	m := make(map[string]interface{})
	m["Cart"] = cart
	err := util.RenderWithAccountAndCommonTem(w, r, m, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 将商品加到购物车
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	itemId := util.GetParam(r, "workingItemId")[0]
	cart := util.GetCartFromSession(w, r, func(cart *domain.Cart) {
		if cart != nil {
			if _, ok := cart.ContainItem(itemId); ok {
				cart.IncrementQuantityByItemId(itemId)
			} else {
				item, err := service.GetItem(itemId)
				isInStock := service.IsItemInStock(itemId)
				if err != nil {
					panic(err)
				}
				cart.AddItem(*item, isInStock)
			}
		}
	})
	// render
	m := make(map[string]interface{})
	m["Cart"] = cart
	err := util.RenderWithAccountAndCommonTem(w, r, m, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 从购物车移除商品
func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	itemId := util.GetParam(r, "workingItemId")[0]
	cart := util.GetCartFromSession(w, r, func(cart *domain.Cart) {
		if cart != nil {
			cart.RemoveItemById(itemId)
		}
	})

	// render
	m := make(map[string]interface{})
	m["Cart"] = cart
	err := util.RenderWithAccountAndCommonTem(w, r, m, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
