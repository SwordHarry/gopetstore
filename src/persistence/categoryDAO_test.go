package persistence

import (
	"fmt"
	"testing"
)

func TestCategoryDAO(t *testing.T) {
	result, err := GetCategoryList()

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(result)
}
