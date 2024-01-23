package main

import (
	"github.com/Bakarseck/orm-go"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	orm.Model
	Username string `orm-go:"NOT NULL UNIQUE"`
	Email    string `orm-go:"NOT NULL UNIQUE"`
}


func main() {
	user := User{}
	orm := orm.NewORM()
	orm.InitDB("mydb.db")
	orm.AutoMigrate(user)
}