package domain

type Item struct {
	ItemId        string
	ProductId     string
	ListPrice     float64
	UnitCost      float64
	SupplierId    int
	Status        string
	AttributeList [5]string
	Product       *Product
	Quantity      int
}

func (i *Item) String() string {
	return i.ItemId
}
