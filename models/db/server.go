package db

import (
	// format
	"fmt"

	// ORMapper
	gorm "github.com/jinzhu/gorm"
	// postgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DbAccess DBアクセサ
var DbAccess *gorm.DB

// init DB初期化
func init() {

	// DB Open
	open()
}

// open DB Open
func open() {

	var err error
	connect := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=sample sslmode=disable"
	DbAccess, err = gorm.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	// 詳細なログを表示
	DbAccess.LogMode(true)
	fmt.Println("db connected: ", &DbAccess)
}

// Close DB Close
func Close() {
	defer DbAccess.Close()
}
