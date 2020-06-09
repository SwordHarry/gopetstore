package domain

type Item struct {
	ItemId        string
	ProductId     string
	ListPrice     float32
	UnitCost      float32
	SupplierId    int
	Status        string
	AttributeList [5]string
	Product
	Quantity int
}

func (i *Item) String() string {
	return i.ItemId
}

func NewItem(
	itemId, productId, status string,
	listPrice, unitCost float32,
	supplierId, quantity int,
	attributeList [5]string,
	product Product,
) Item {
	return Item{
		ItemId:        itemId,
		ProductId:     productId,
		ListPrice:     listPrice,
		UnitCost:      unitCost,
		SupplierId:    supplierId,
		Status:        status,
		AttributeList: attributeList,
		Product:       product,
		Quantity:      quantity,
	}
}
