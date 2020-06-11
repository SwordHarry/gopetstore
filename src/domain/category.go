package domain

type Category struct {
	CategoryId  string
	Name        string
	Description string
}

func (c *Category) String() string {
	return c.CategoryId
}
