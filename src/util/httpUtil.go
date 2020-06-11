package util

import (
	"net/http"
)

// about request
// 获取 url 中的 参数
func GetParam(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}
