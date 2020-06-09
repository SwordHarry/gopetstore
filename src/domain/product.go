package domain

type Product struct {
	ProductId   string
	CategoryId  string
	Name        string
	Description string
}

func (p *Product) String() string {
	return p.ProductId
}

func NewProduct(productId, categoryId, name, description string) Product {
	return Product{
		ProductId:   productId,
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
	}
}
