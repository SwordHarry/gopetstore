package util

import (
	"database/sql"
	// 驱动需要进行隐式导入
	_ "github.com/go-sql-driver/mysql"
)

const userName = "root"
const password = "root"
const dbName = "gopetstore"
const driverName = "mysql"
const charset = "charset=utf8"
const local = "loc=Local"
const tcpPort = "@tcp(localhost:3306)/"

// 连接数据库 mysql
func GetConnection() (*sql.DB, error) {
	dataSourceName := userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local
	db, err := sql.Open(driverName, dataSourceName) //对应数据库的用户名和密码以及数据库名

	return db, err
}
