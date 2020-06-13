package config

import "path/filepath"

// dir name
const (
	Front   = "front"
	Web     = "web"
	Catalog = "catalog"
	Common  = "common"
	Cart    = "cart"
	Account = "account"
	Order   = "order"
)

var CommonPath = filepath.Join(Front, Web, Common, "common.html")
