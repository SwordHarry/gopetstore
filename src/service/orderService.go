package service

import (
	"gopetstore/src/domain"
	"gopetstore/src/persistence"
	"log"
)

const orderNum = "ordernum"

// get order by order id
func GetOrderByOrderId(orderId int) (*domain.Order, error) {
	o, err := persistence.GetOrderByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	o.LineItems, err = persistence.GetLineItemsByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	for _, li := range o.LineItems {
		item, err := persistence.GetItem(li.ItemId)
		if err != nil {
			log.Printf("service GetOrderByOrderId GetItem error: %v", err.Error())
			continue
		}
		item.Quantity, err = persistence.GetInventoryQuantity(li.ItemId)
		if err != nil {
			log.Printf("service GetOrderByOrderId GetInventoryQuantity error: %v", err.Error())
			continue
		}
		li.Item = item
	}
	return o, nil
}

// get all orders by user name
func GetOrdersByUserName(userName string) ([]*domain.Order, error) {
	return persistence.GetOrdersByUserName(userName)
}

// insert order
func InsertOrder(o *domain.Order) error {
	orderId, err := getNextId(orderNum)
	if err != nil {
		return err
	}
	o.OrderId = orderId
	return persistence.InsertOrder(o)
}

// update the sequence and next id
func getNextId(name string) (int, error) {
	s, err := persistence.GetSequence(name)
	if err != nil {
		return -1, err
	}
	s.NextId++
	err = persistence.UpdateSequence(s)
	if err != nil {
		return -1, err
	}
	return s.NextId, nil
}
