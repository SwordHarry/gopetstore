package util

import (
	"testing"
)

func TestDBUtil(t *testing.T) {
	d, err := GetConnection()
	defer d.Close()
	if err != nil {
		t.Errorf("数据库链接错误: %v", err.Error())
	}
}
