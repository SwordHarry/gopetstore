package domain

// 购物项，单价和数量
type CartItem struct {
	Item     *Item
	Quantity int
	InStock  bool
	Total    float32
}

// 数量增加
func (ci *CartItem) incrementQuantity() {
	ci.Quantity++
	ci.calculateTotal()
}

// 计算总价: 单价 * 数量
func (ci *CartItem) calculateTotal() {
	if ci.Item != nil && ci.Item.ListPrice != 0 {
		ci.Total = ci.Item.ListPrice * float32(ci.Quantity)
	}
}
