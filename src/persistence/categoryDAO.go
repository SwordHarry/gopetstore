package persistence

import (
	"database/sql"
	"errors"
	"gopetstore/src/domain"
	"log"
)

const getCategoryListSQL = "SELECT CATID AS categoryId,NAME,DESCN AS description FROM CATEGORY"
const getCategoryByIdSQL = "SELECT CATID AS categoryId,NAME,DESCN AS description FROM CATEGORY WHERE CATID = ?"

// 代码封装：通用 scan category 逻辑
func scanCategory(r *sql.Rows) (*domain.Category, error) {
	var categoryId, name, description string
	err := r.Scan(&categoryId, &name, &description)
	if err != nil {
		return nil, err
	}
	return &domain.Category{
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
	}, nil
}

// 获取所有的 category
func GetCategoryList() ([]domain.Category, error) {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	var result []domain.Category
	if err != nil {
		return result, err
	}
	r, err := d.Query(getCategoryListSQL)
	if err != nil {
		return result, err
	}
	for r.Next() {
		c, err := scanCategory(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		result = append(result, *c)
	}
	return result, nil
}

// 通过 id 获取指定的 category
func GetCategory(categoryId string) (*domain.Category, error) {
	d, err := getConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getCategoryByIdSQL, categoryId)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		c, err := scanCategory(r)
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	return nil, errors.New("can not find a category by this id")
}
