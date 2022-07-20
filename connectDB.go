package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func connectDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:hanhuquynh@tcp(127.0.0.1:3306)/demo?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	return engine
}
