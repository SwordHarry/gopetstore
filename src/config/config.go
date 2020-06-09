package config

import "path/filepath"

const Front = "front"
const Web = "web"
const Catalog = "catalog"
const COMMON = "common"

var Common = filepath.Join(Front, Web, COMMON, "common.html")
