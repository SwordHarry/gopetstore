package persistence

import (
	"database/sql"
	"errors"
	"gopetstore/src/domain"
	"log"
)

const getItemByIdSQL = `select I.ITEMID,LISTPRICE,UNITCOST,SUPPLIER AS supplierId,I.PRODUCTID AS productId,
NAME AS productName,DESCN AS productDescription,CATEGORY AS CategoryId,STATUS,
IFNULL(ATTR1, "") AS attribute1,IFNULL(ATTR2, "") AS attribute2,IFNULL(ATTR3, "") AS attribute3,
IFNULL(ATTR4, "") AS attribute4,IFNULL(ATTR5, "") AS attribute5,QTY AS quantity from ITEM I, INVENTORY V, PRODUCT P 
where P.PRODUCTID = I.PRODUCTID and I.ITEMID = V.ITEMID and I.ITEMID=?`

const getItemListByProductIdSQL = `SELECT I.ITEMID,LISTPRICE,UNITCOST,SUPPLIER AS supplierId,I.PRODUCTID AS productId,
NAME AS productName,DESCN AS productDescription,CATEGORY AS categoryId,STATUS,
IFNULL(ATTR1, "") AS attribute1,IFNULL(ATTR2, "") AS attribute2,IFNULL(ATTR3, "") AS attribute3,
IFNULL(ATTR4, "") AS attribute4,IFNULL(ATTR5, "") AS attribute5 FROM ITEM I, PRODUCT P 
WHERE P.PRODUCTID = I.PRODUCTID AND I.PRODUCTID = ?`
const getInventoryByItemIdSQL = `SELECT QTY AS QUANTITY FROM INVENTORY WHERE ITEMID = ?`
const updateInventoryByItemIdSQl = `UPDATE INVENTORY SET QTY = QTY - ? WHERE ITEMID = ?`

// scan item, hasQuantity is for the num of param is not equal
func scanItem(r *sql.Rows, hasQuantity bool) (*domain.Item, error) {
	var itemId, productId, status string
	var listPrice, unitCost float32
	var supplierId, quantity int
	var attributeList [5]string
	var pName, pDescription, pCategoryId string
	var err error
	if hasQuantity {
		err = r.Scan(&itemId, &listPrice, &unitCost, &supplierId, &productId,
			&pName, &pDescription, &pCategoryId, &status,
			&attributeList[0], &attributeList[1], &attributeList[2], &attributeList[3], &attributeList[4], &quantity)
	} else {
		err = r.Scan(&itemId, &listPrice, &unitCost, &supplierId, &productId,
			&pName, &pDescription, &pCategoryId, &status,
			&attributeList[0], &attributeList[1], &attributeList[2], &attributeList[3], &attributeList[4])
	}
	if err != nil {
		return nil, err
	}
	p := domain.Product{
		ProductId:   productId,
		CategoryId:  pCategoryId,
		Name:        pName,
		Description: pDescription,
	}

	i := domain.Item{
		ItemId:        itemId,
		ProductId:     productId,
		ListPrice:     listPrice,
		UnitCost:      unitCost,
		SupplierId:    supplierId,
		Status:        status,
		AttributeList: attributeList,
		Product:       p,
		Quantity:      quantity,
	}
	return &i, nil
}

// get item by item id
func GetItem(itemId string) (*domain.Item, error) {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getItemByIdSQL, itemId)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		i, err := scanItem(r, true)
		if err != nil {
			return nil, err
		}
		return i, nil
	}
	return nil, errors.New("can not find the item by this id")
}

// get all items by product id
func GetItemListByProduct(productId string) ([]domain.Item, error) {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	var result []domain.Item
	if err != nil {
		return result, err
	}
	r, err := d.Query(getItemListByProductIdSQL, productId)
	if err != nil {
		return result, err
	}
	for r.Next() {
		p, err := scanItem(r, false)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		result = append(result, *p)
	}
	return result, nil
}

// get inventory by item id
func GetInventoryQuantity(itemId string) (int, error) {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return -1, err
	}
	r, err := d.Query(getInventoryByItemIdSQL, itemId)
	if err != nil {
		return -1, err
	}
	if r.Next() {
		var result int
		err := r.Scan(&result)
		if err != nil {
			return -1, err
		}
		return result, nil
	}
	return -1, errors.New("can not find the inventory by this id")
}

// update inventory by item id
func UpdateInventoryQuantity(itemId string, increment int) error {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return err
	}
	r, err := d.Exec(updateInventoryByItemIdSQl, increment, itemId)
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if rowNum > 0 && err == nil {
		return nil
	}
	return err
}
