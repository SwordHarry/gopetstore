package domain

import (
	"encoding/gob"
)

type Product struct {
	ProductId   string
	CategoryId  string
	Name        string
	Description string
}

func (p *Product) String() string {
	return p.ProductId
}

// 序列化注册 product，用于 session 存储
func init() {
	gob.Register(&Product{})
}
