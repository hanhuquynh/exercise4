package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var engine = connectDB()

func createTable() {
	engine.Sync2(new(User), new(Point))
}

func insertUser(user *User, point *Point) {
	_, err := engine.Table("user").Insert(user)

	if err != nil {
		log.Fatal(err)
	}

	_, err = engine.Table("point").Insert(point)

	if err != nil {
		log.Fatal(err)
	}
}

func listUser() {
	var user []User

	err := engine.Find(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func userById(id string) {
	var user User

	has, err := engine.Exist(&User{
		Id: id,
	})

	if err != nil {
		log.Fatal(err)
	}

	if has {
		_, err := engine.Where("id = ?", id).Get(&user)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
	} else {
		fmt.Println("ID doesn't exist")
	}

}

func updateUser(user *User, id string) {
	_, err := engine.Where("id = ?", id).Update(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update successful")
}
