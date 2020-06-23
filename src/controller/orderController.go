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

const (
	viewOrderFile    = "viewOrder.html"
	initOrderFile    = "initOrder.html"
	confirmOrderFile = "confirmOrder.html"
	shipFormFile     = "shipForm.html"
	listOrdersFile   = "listOrders.html"
)

var (
	viewOrderPath    = filepath.Join(config.Front, config.Web, config.Order, viewOrderFile)
	initOrderPath    = filepath.Join(config.Front, config.Web, config.Order, initOrderFile)
	confirmOrderPath = filepath.Join(config.Front, config.Web, config.Order, confirmOrderFile)
	shipFormPath     = filepath.Join(config.Front, config.Web, config.Order, shipFormFile)
	listOrderPath    = filepath.Join(config.Front, config.Web, config.Order, listOrdersFile)
)

// [ViewInitOrder] -> initOrder -> [ConfirmOrderStep1] -> shipForm -> [ConfirmShip] ->
// confirmOrder -> [ConfirmOrderStep2] -> viewOrder

// ViewOrderList -> viewOrder

// render init order
func ViewInitOrder(w http.ResponseWriter, r *http.Request) {
	account := util.GetAccountFromSession(r)
	if account != nil {
		cart := util.GetCartFromSession(w, r, nil)
		if cart != nil {
			o := domain.NewOrder(account, cart)
			s, err := util.GetSession(r)
			if err != nil {
				log.Printf("ViewInitOrder GetSession error: %v", err.Error())
			}
			if s != nil {
				err = s.Save(config.OrderKey, o, w, r)
				if err != nil {
					log.Printf("ViewInitOrder GetSession error: %v", err.Error())
				}
				m := make(map[string]interface{})
				m["Order"] = o
				m["CreditCardTypes"] = []string{o.CardType}
				err = util.RenderWithAccountAndCommonTem(w, r, m, initOrderPath)
				if err != nil {
					log.Printf("ViewInitOrder RenderWithAccountAndCommonTem error: %v", err.Error())
				}
			} else {
				log.Print("ViewInitOrder session is nil")
			}
		} else {
			log.Print("ViewInitOrder cart is nil")
		}
	} else {
		// 跳转到登录页面
		ViewLoginOrPostLogin(w, r)
	}
}

// press confirm button step 1
func ConfirmOrderStep1(w http.ResponseWriter, r *http.Request) {
	// get order from session
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ConfirmOrderStep1 GetSession error: %v", err.Error())
	}
	if s != nil {
		re, _ := s.Get(config.OrderKey)
		o := re.(*domain.Order)
		// post parse form
		err := r.ParseForm()
		if err != nil {
			log.Printf("ConfirmOrderStep1 ParseForm error: %v", err.Error())
			return
		}
		o.CardType = r.FormValue("cardType")
		o.CreditCard = r.FormValue("creditCard")
		o.ExpiryDate = r.FormValue("expiryDate")
		o.BillToFirstName = r.FormValue("firstName")
		o.BillToLastName = r.FormValue("lastName")
		o.BillAddress1 = r.FormValue("address1")
		o.BillAddress2 = r.FormValue("address2")
		o.BillCity = r.FormValue("city")
		o.BillState = r.FormValue("state")
		o.BillZip = r.FormValue("zip")
		o.BillCountry = r.FormValue("country")
		m := make(map[string]interface{})
		m["Order"] = o
		if len(r.FormValue("shippingAddressRequired")) > 0 {
			// view shipForm
			err := util.RenderWithAccountAndCommonTem(w, r, m, shipFormPath)
			if err != nil {
				log.Printf("ConfirmOrderStep1 RenderWithAccountAndCommonTem error: %v", err.Error())
			}
		} else {
			// view confirmOrder
			err := util.RenderWithAccountAndCommonTem(w, r, m, confirmOrderPath)
			if err != nil {
				log.Printf("ConfirmOrderStep1 RenderWithAccountAndCommonTem error: %v", err.Error())
			}
		}
	}
}

// confirm ship
func ConfirmShip(w http.ResponseWriter, r *http.Request) {
	// get order from session
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ConfirmShip GetSession error: %v", err.Error())
	}
	if s != nil {
		re, _ := s.Get(config.OrderKey)
		o := re.(*domain.Order)
		err := r.ParseForm()
		if err != nil {
			log.Printf("ConfirmShip ParseForm error: %v", err.Error())
			return
		}
		o.ShipToFirstName = r.FormValue("shipToFirstName")
		o.ShipToLastName = r.FormValue("shipToLastName")
		o.ShipAddress1 = r.FormValue("shipAddress1")
		o.ShipAddress2 = r.FormValue("shipAddress2")
		o.ShipCity = r.FormValue("shipCity")
		o.ShipState = r.FormValue("shipState")
		o.ShipZip = r.FormValue("shipZip")
		o.ShipCountry = r.FormValue("shipCountry")
		m := make(map[string]interface{})
		m["Order"] = o
		err = util.RenderWithAccountAndCommonTem(w, r, m, confirmOrderPath)
		if err != nil {
			log.Printf("ConfirmShip RenderWithAccountAndCommonTem error: %v", err.Error())
		}
	}
}

// create the final order
func ConfirmOrderStep2(w http.ResponseWriter, r *http.Request) {
	// get order from session
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ConfirmShip GetSession error: %v", err.Error())
	}
	if s != nil {
		re, _ := s.Get(config.OrderKey)
		o := re.(*domain.Order)
		err := service.InsertOrder(o)
		if err != nil {
			log.Printf("ConfirmOrderStep2 InsertOrder error: %v", err.Error())
			return
		}
		// 清空购物车

		err = s.Del(config.CartKey, w, r)
		if err != nil {
			log.Printf("ConfirmOrderStep2 session del cart error: %v", err.Error())
		}
		m := map[string]interface{}{
			"Order": o,
		}
		err = util.RenderWithAccountAndCommonTem(w, r, m, viewOrderPath)
		if err != nil {
			log.Printf("ConfirmOrderStep2 RenderWithAccountAndCommonTem error: %v", err.Error())
		}
	}
}

// list orders
func ListOrders(w http.ResponseWriter, r *http.Request) {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ListOrders GetSession error: %v", err.Error())
	}
	if s != nil {
		re, ok := s.Get(config.AccountKey)
		if ok {
			a, ok := re.(*domain.Account)
			if ok {
				orders, err := service.GetOrdersByUserName(a.UserName)
				if err != nil {
					log.Printf("ListOrders GetOrdersByUserName error: %v", err.Error())
				}
				m := map[string]interface{}{
					"OrderList": orders,
				}
				err = util.RenderWithAccountAndCommonTem(w, r, m, listOrderPath)
				if err != nil {
					log.Printf("ListOrders RenderWithAccountAndCommonTem error: %v", err.Error())
				}
			}
		}
	}
}

// check order
func CheckOrder(w http.ResponseWriter, r *http.Request) {
	orderIdStr := util.GetParam(r, "orderId")[0]
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		log.Printf("CheckOrder error: %v", err.Error())
	}

	o, err := service.GetOrderByOrderId(orderId)
	m := map[string]interface{}{
		"Order": o,
	}
	err = util.RenderWithAccountAndCommonTem(w, r, m, viewOrderPath)
	if err != nil {
		log.Printf("CheckOrder RenderWithAccountAndCommonTem error: %v", err.Error())
	}
}
