package route

/**
将路由改为 map 映射的配置
*/

import (
	"gopetstore/src/controller"
	"net/http"
)

// 路由映射注册表
var route = map[string]http.HandlerFunc{
	// view
	"/main":          controller.ViewMain,
	"/viewCategory":  controller.ViewCategory,
	"/viewProduct":   controller.ViewProduct,
	"/viewItem":      controller.ViewItem,
	"/searchProduct": controller.SearchProduct,
	// cart
	"/addItemToCart":      controller.AddItemToCart,
	"/viewCart":           controller.ViewCart,
	"/removeItemFromCart": controller.RemoveItemFromCart,
	// account
	"/login":       controller.ViewLoginOrPostLogin,
	"/register":    controller.ViewRegister,
	"/signOut":     controller.SignOut,
	"/editAccount": controller.ViewEditAccount,
	"/newAccount":  controller.NewAccount,
	"/confirmEdit": controller.ConfirmEdit,
}

// 注册路由
func RegisterRoute() {
	for k, v := range route {
		http.HandleFunc(k, v)
	}
}
