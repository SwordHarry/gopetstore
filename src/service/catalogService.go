package service

/**
catalog service 层，整合DAO层 model 逻辑
category 和 product 的 DAO层
*/

import (
	"gopetstore/src/domain"
	"gopetstore/src/persistence"
	"log"
)

// get category by id
func GetCategory(categoryId string) (*domain.Category, error) {
	return persistence.GetCategory(categoryId)
}

// get all categories
func GetCategoryList() ([]domain.Category, error) {
	return persistence.GetCategoryList()
}

// get product by id
func GetProduct(productId string) (*domain.Product, error) {
	return persistence.GetProduct(productId)
}

// get all products by category id
func GetProductList(categoryId string) ([]domain.Product, error) {
	return persistence.GetProductListByCategory(categoryId)
}

// get products by keyword
func SearchProductList(keyword string) ([]domain.Product, error) {
	return persistence.SearchProductList(keyword)
}

// get items by product id
func GetItemListByProduct(productId string) ([]domain.Item, error) {
	return persistence.GetItemListByProduct(productId)
}

// get item by item id
func GetItem(itemId string) (*domain.Item, error) {
	return persistence.GetItem(itemId)
}

// is item in stock
func IsItemInStock(itemId string) bool {
	flag, err := persistence.GetInventoryQuantity(itemId)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return false
	}
	return flag > 0
}
