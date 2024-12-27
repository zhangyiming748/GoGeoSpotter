package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"xorm.io/xorm"
)

var (
	useMysql bool
	engine   *xorm.Engine
)

// CREATE DATABASE `tdl` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
func SetMysql() {
	var err error
	user := "root"
	password := "163453"
	host := os.Getenv("MYSQL_HOST")
	//host := "192.168.1.9"
	port := os.Getenv("MYSQL_PORT")
	//port := "3306"
	dbName := "amap"
	charset := "utf8"
	engine, err = xorm.NewEngine("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset)
	//engine, err = xorm.NewEngine("mysql", "root:163453@tcp(192.168.1.9:3306)/amap?charset=utf8")
	if err != nil {
		useMysql = false
	} else {
		useMysql = true
		log.Printf("连接数据库成功:%v\n", engine)
	}
}

func GetMysql() *xorm.Engine {
	return engine
}

func UseMysql() bool {
	return useMysql
}
