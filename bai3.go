package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func insert100User() {
	time := time.Now().UnixMilli()
	for i := 1; i <= 100; i++ {
		insertUser(&User{
			Id:         "u" + strconv.Itoa(i),
			Name:       "User " + strconv.Itoa(i),
			Birth:      2001,
			Created:    time,
			Updated_at: time,
		}, &Point{
			User_id:    "u" + strconv.Itoa(i),
			Points:     10,
			Max_points: 10,
		})
	}
}

func scanTableUser() {
	ch := make(chan *User, 10)

	defer close(ch)

	user := User{}
	rows, err := engine.Rows(&user)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	bean := new(User)

	count, err := engine.Count(&user)

	if err != nil {
		log.Fatal(err)
	}

	intCount := int(count)

	for i := 0; i < 2; i++ {
		go worker(ch, intCount)
	}

	for rows.Next() {
		err := rows.Scan(bean)

		if err != nil {
			log.Fatal(err)
		}

		ch <- bean
	}
}

func worker(u chan *User, count int) {
	for i := 1; i <= count; i++ {
		data := <-u
		fmt.Printf("%v - %v - %v\n", i, data.Id, data.Name)
	}
}
