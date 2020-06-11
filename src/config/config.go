package config

import "path/filepath"

const Front = "front"
const Web = "web"
const Catalog = "catalog"
const COMMON = "common"
const Cart = "cart"

var CommonPath = filepath.Join(Front, Web, COMMON, "common.html")
