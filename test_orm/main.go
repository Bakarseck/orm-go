package main

import (
	"fmt"

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
	User_id      int64 `orm-go:"NOT NULL FOREIGN_KEY:User:Id"`
}

func NewUser(name, email string) User {
	return User{
		Username: name,
		Email:    email,
	}
}

func NewProduit(name string, p int64) Produit {
	return Produit{
		Name_produit: name,
		Prix:         p,
	}
}

func main() {
	user := User{}
	produit := Produit{}

	orm := orm.NewORM()
	orm.InitDB("test.db")
	orm.AutoMigrate(user, produit)

	users := orm.Scan(User{}, "Id", "CreatedAt", "Email", "Username").([]User)

	for _, user := range users {
		fmt.Println("User: ", user)
	}

}
