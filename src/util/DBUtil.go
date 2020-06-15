package util

import (
	"database/sql"
	"errors"
	"log"
	// 驱动需要进行隐式导入
	_ "github.com/go-sql-driver/mysql"
)

const (
	userName   = "root"
	password   = "root"
	dbName     = "gopetstore"
	driverName = "mysql"
	charset    = "charset=utf8"
	local      = "loc=Local"
	tcpPort    = "@tcp(localhost:3306)/"
	parseTime  = "parseTime=true" // 用以解析 数据库 中的 date 类型，否则会解析成 []uint8 不能隐式转为 string
)

// 连接数据库 mysql
func GetConnection() (*sql.DB, error) {
	dataSourceName := userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local + "&" + parseTime
	db, err := sql.Open(driverName, dataSourceName) //对应数据库的用户名和密码以及数据库名

	return db, err
}

// 发现 更新和插入 语句的逻辑和返回值几乎一致，这里进行再封装
func InsertOrUpdate(SQL string, errStr string, param ...interface{}) error {
	d, err := GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return err
	}
	log.Print(param)
	var r sql.Result
	if len(param) > 0 {
		r, err = d.Exec(SQL, param...)
	} else {
		r, err = d.Exec(SQL)
	}
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum > 0 {
		return nil
	}
	log.Print(rowNum)
	return errors.New(errStr)
}
