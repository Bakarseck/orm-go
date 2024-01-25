package main

import (
	orm "github.com/Bakarseck/orm-go/ORM"
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
	Prix         int64
}

type Comment struct {
	orm.Model
	Name_produit string `orm-go:"NOT NULL"`
	Prix         int64
}

type Post struct {
	orm.Model
	Title   string `orm-go:"NOT NULL"`
	Content string `orm-go:"NOT NULL"`
	UserId  int64  `orm-go:"FOREIGN_KEY:User"`
}

func main() {
	user := User{}
	produit := Produit{}
	c := Comment{}

	p := Post{}

	orm := orm.NewORM()
	orm.InitDB("mydb.db")
	orm.AutoMigrate(user, produit, c, p)
}
