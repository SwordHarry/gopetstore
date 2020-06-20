package persistence

import (
	"database/sql"
	"gopetstore/src/domain"
	"gopetstore/src/util"
	"log"
)

const (
	getLineItemsByOrderIdSQL = `SELECT ORDERID, LINENUM AS lineNumber, ITEMID, QUANTITY, UNITPRICE FROM LINEITEM WHERE ORDERID = ?`
	insertLineItemSQL        = `INSERT INTO LINEITEM (ORDERID, LINENUM, ITEMID, QUANTITY, UNITPRICE) VALUES (?, ?, ?, ?, ?)`
)

// scan line item
func scanLineItem(r *sql.Rows) (*domain.LineItem, error) {
	var orderId, lineNumber, quantity int
	var itemId string
	var unitPrice float32
	err := r.Scan(&orderId, &lineNumber, &itemId, &quantity, &unitPrice)
	if err != nil {
		return nil, err
	}
	li := &domain.LineItem{
		OrderId:    orderId,
		LineNumber: lineNumber,
		Quantity:   quantity,
		ItemId:     itemId,
		UnitPrice:  unitPrice,
	}
	return li, nil
}

// get line item by order id
func GetLineItemsByOrderId(orderId int) ([]*domain.LineItem, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getLineItemsByOrderIdSQL, orderId)
	defer func() {
		if r != nil {
			_ = r.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	var result []*domain.LineItem
	for r.Next() {
		li, err := scanLineItem(r)
		if err != nil {
			log.Printf("scanLineItem error: %v", err.Error())
			continue
		}
		result = append(result, li)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}
