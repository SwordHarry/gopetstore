package persistence

import (
	"database/sql"
	"errors"
	"gopetstore/src/domain"
	"gopetstore/src/util"
	"log"
)

const getProductListByCategorySQL = "SELECT PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId FROM PRODUCT WHERE CATEGORY = ?"
const getProductByIdSQL = "SELECT PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId FROM PRODUCT WHERE PRODUCTID = ?"
const getProductListByKeyword = "select PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId from PRODUCT WHERE lower(name) like ?"

// 代码封装：通用 scan product 逻辑
func scanProduct(r *sql.Rows) (*domain.Product, error) {
	var productId, name, categoryId, description string
	err := r.Scan(&productId, &name, &description, &categoryId)
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ProductId:   productId,
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
	}, nil
}

// 通过 category 的 id 获取之下的 product
func GetProductListByCategory(categoryId string) ([]domain.Product, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	var result []domain.Product
	if err != nil {
		return result, err
	}
	r, err := d.Query(getProductListByCategorySQL, categoryId)
	if err != nil {
		return result, err
	}
	for r.Next() {
		p, err := scanProduct(r)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		result = append(result, *p)
	}
	defer r.Close()
	return result, nil
}

// 通过 id 获取 product
func GetProduct(productId string) (*domain.Product, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getProductByIdSQL, productId)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		p, err := scanProduct(r)
		if err != nil {
			return nil, err
		}

		return p, nil
	}
	defer r.Close()
	// 这里的逻辑是 r.Next() 中没有值
	return nil, errors.New("can not find a product by this id")
}

// 通过关键字获取 product
func SearchProductList(keyword string) ([]domain.Product, error) {
	var result []domain.Product
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return result, err
	}
	r, err := d.Query(getProductListByKeyword, keyword)
	defer func() {
		if r != nil {
			_ = r.Close()
		}
	}()
	if err != nil {
		return result, err
	}
	for r.Next() {
		p, err := scanProduct(r)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		result = append(result, *p)
	}
	return result, nil
}
