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

// 将商品加到购物车
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	itemId := util.GetParam(r, "workingItemId")[0]
	cart := getCartFromSession(w, r, func(cart *domain.Cart) {
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

	err := util.RenderWithCommon(w, struct {
		Account interface{}
		Cart    *domain.Cart
	}{
		Account: nil,
		Cart:    cart,
	}, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 跳转到购物车页面
func ViewCart(w http.ResponseWriter, r *http.Request) {
	cart := getCartFromSession(w, r, func(cart *domain.Cart) {
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
	err := util.RenderWithCommon(w, struct {
		Account interface{}
		Cart    *domain.Cart
	}{
		Account: nil,
		Cart:    cart,
	}, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 从购物车移除商品
func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	itemId := util.GetParam(r, "workingItemId")[0]
	cart := getCartFromSession(w, r, func(cart *domain.Cart) {
		if cart != nil {
			cart.RemoveItemById(itemId)
		}
	})

	err := util.RenderWithCommon(w, struct {
		Account interface{}
		Cart    *domain.Cart
	}{
		Account: nil,
		Cart:    cart,
	}, cartPath)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}

// 从session中获取cart
func getCartFromSession(w http.ResponseWriter, r *http.Request, callback func(cart *domain.Cart)) *domain.Cart {
	// 使用 session 存储 cart 购物车
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("session error for getSession: %v", err.Error())
	}
	var cart *domain.Cart
	// 成功生成 session
	if s != nil {
		c, ok := s.Get("cart")
		if !ok {
			// 初始化 购物车
			c = domain.NewCart()
		}
		// 调用回调对cart 进行操作
		cart, ok = c.(*domain.Cart)
		if ok && callback != nil {
			callback(cart)
		}
		// 将新的购物车进行存储覆盖
		err := s.Save("cart", c, w, r)
		if err != nil {
			log.Printf("session error for Save: %v", err.Error())
		}
	}
	return cart
}
