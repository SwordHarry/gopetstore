package persistence

import (
	"database/sql"
	"errors"
	"gopetstore/src/domain"
	"gopetstore/src/util"
	"log"
	"time"
)

const (
	getOrderByOrderIdSQL = `select BILLADDR1 AS billAddress1,BILLADDR2 AS billAddress2,BILLCITY,BILLCOUNTRY,BILLSTATE,BILLTOFIRSTNAME,BILLTOLASTNAME,BILLZIP,
SHIPADDR1 AS shipAddress1,SHIPADDR2 AS shipAddress2,SHIPCITY,SHIPCOUNTRY,SHIPSTATE,SHIPTOFIRSTNAME,SHIPTOLASTNAME,SHIPZIP,CARDTYPE,COURIER,CREDITCARD,
EXPRDATE AS expiryDate,LOCALE,ORDERDATE,ORDERS.ORDERID,TOTALPRICE,USERID AS username,STATUS FROM ORDERS, ORDERSTATUS 
WHERE ORDERS.ORDERID = ? AND ORDERS.ORDERID = ORDERSTATUS.ORDERID`
	getOrdersByUsernameSQL = `SELECT BILLADDR1 AS billAddress1, BILLADDR2 AS billAddress2, BILLCITY, BILLCOUNTRY, BILLSTATE, BILLTOFIRSTNAME, BILLTOLASTNAME, BILLZIP,
SHIPADDR1 AS shipAddress1, SHIPADDR2 AS shipAddress2, SHIPCITY, SHIPCOUNTRY, SHIPSTATE, SHIPTOFIRSTNAME, SHIPTOLASTNAME, SHIPZIP, CARDTYPE, COURIER, CREDITCARD, EXPRDATE AS expiryDate,LOCALE,
ORDERDATE, ORDERS.ORDERID, TOTALPRICE, USERID AS username,STATUS FROM ORDERS, ORDERSTATUS WHERE ORDERS.USERID = ? AND ORDERS.ORDERID = ORDERSTATUS.ORDERID ORDER BY ORDERDATE`
	insertOrderSQL = `INSERT INTO ORDERS (ORDERID, USERID, ORDERDATE, SHIPADDR1, SHIPADDR2, SHIPCITY, SHIPSTATE, SHIPZIP, SHIPCOUNTRY,
BILLADDR1, BILLADDR2, BILLCITY, BILLSTATE, BILLZIP, BILLCOUNTRY, COURIER, TOTALPRICE, BILLTOFIRSTNAME, BILLTOLASTNAME, SHIPTOFIRSTNAME, SHIPTOLASTNAME, CREDITCARD, EXPRDATE, CARDTYPE, LOCALE) 
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	insertOrderStatusSQL = `INSERT INTO ORDERSTATUS (ORDERID, LINENUM, TIMESTAMP, STATUS) VALUES (?, ?, ?, ?)`
)

// scan order
func scanOrder(r *sql.Rows) (*domain.Order, error) {
	var billAddr1, billAddr2, billCity, billCountry, billState, billToFirstName, billToLastName, billZip string
	var shipAddr1, shipAddr2, shipCity, shipCountry, shipState, shipFirstName, shipLastName, shipZip string
	var cardType, courier, creditCard string
	var expiryDate, locale, userName, status string
	var totalPrice float32
	var orderDate time.Time
	var orderId int
	err := r.Scan(&billAddr1, &billAddr2, &billCity, &billCountry, &billState, &billToFirstName, &billToLastName, &billZip,
		&shipAddr1, &shipAddr2, &shipCity, &shipCountry, &shipState, &shipFirstName, &shipLastName, &shipZip,
		&cardType, &courier, &creditCard, &expiryDate, &locale, &orderDate, &orderId, &totalPrice, &userName, &status)
	if err != nil {
		return nil, err
	}
	return &domain.Order{
		OrderId:         orderId,
		OrderDate:       orderDate,
		UserName:        userName,
		ShipAddress1:    shipAddr1,
		ShipAddress2:    shipAddr2,
		ShipCity:        shipCity,
		ShipState:       shipState,
		ShipCountry:     shipCountry,
		ShipToFirstName: shipFirstName,
		ShipToLastName:  shipLastName,
		BillAddress1:    billAddr1,
		BillAddress2:    billAddr2,
		BillCity:        billCity,
		BillZip:         billZip,
		BillCountry:     billCountry,
		BillToFirstName: billToFirstName,
		BillToLastName:  billToLastName,
		Courier:         courier,
		CreditCard:      creditCard,
		CardType:        cardType,
		TotalPrice:      totalPrice,
		ExpiryDate:      expiryDate,
		Locale:          locale,
		Status:          status,
	}, nil
}

// get order by order id
func GetOrderByOrderId(orderId int) (*domain.Order, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getOrderByOrderIdSQL, orderId)
	defer func() {
		if r != nil {
			_ = r.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	if r.Next() {
		order, err := scanOrder(r)
		if err != nil {
			return nil, err
		}
		order.OrderId = orderId
		return order, nil
	}

	return nil, errors.New("can not get a order by this orderId")
}

// get all orders by user name
func GetOrdersByUserName(userName string) ([]*domain.Order, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getOrdersByUsernameSQL, userName)
	defer func() {
		if r != nil {
			_ = r.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	var result []*domain.Order
	for r.Next() {
		order, err := scanOrder(r)
		if err != nil {
			log.Printf("GetOrdersByUserName scanOrder error: %v for userName: %v", err.Error(), userName)
			continue
		}
		result = append(result, order)
	}
	return result, nil
}

// insert order
func InsertOrder(o *domain.Order) error {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return err
	}
	r, err := d.Exec(insertOrderSQL, o.OrderId, o.UserName, o.OrderDate, o.ShipAddress1, o.ShipAddress2, o.ShipCity,
		o.ShipState, o.ShipZip, o.ShipCountry, o.BillAddress1, o.BillAddress2, o.BillCity, o.BillState, o.BillZip,
		o.BillCountry, o.Courier, o.TotalPrice, o.BillToFirstName, o.BillToLastName, o.ShipToFirstName, o.ShipToLastName,
		o.CreditCard, o.ExpiryDate, o.CardType, o.Locale)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}
	return errors.New("insert order failed")
}

// insert order status
func InsertOrderStatus(o *domain.Order) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	r, err := d.Exec(insertOrderStatusSQL, o.OrderId, o.OrderId, o.OrderDate, o.Status)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}
	return errors.New("can not insert order status")
}
