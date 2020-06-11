package config

import "path/filepath"

const Front = "front"
const Web = "web"
const Catalog = "catalog"
const Common = "common"
const Cart = "cart"
const Account = "account"

var CommonPath = filepath.Join(Front, Web, Common, "common.html")
