package main

import (
	"github.com/Bakarseck/orm-go/ORM"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	orm.Model
	Username string `orm-go:"NOT NULL UNIQUE"`
	Email    string `orm-go:"NOT NULL UNIQUE"`
}

type Produit struct {
	orm.Model
	Name_produit string `orm-go:"NOT NULL"`
	Prix int64
}

func main() {
	user := User{}
	produit := Produit{}
	orm := orm.NewORM()
	orm.InitDB("mydb.db")
	orm.AutoMigrate(user, produit)
}
