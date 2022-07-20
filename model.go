package main

type User struct {
	Id         string
	Name       string
	Birth      int64
	Created    int64
	Updated_at int64
}

type Point struct {
	User_id    string
	Points     int64
	Max_points int64
}
