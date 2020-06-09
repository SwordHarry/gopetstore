package domain

type Category struct {
	CategoryId string
	Name string
	Description string
}

func (c *Category) String() string {
	return c.CategoryId
}

func NewCategory(categoryId string,name string,description string) Category  {
	return Category{
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
	}
}
