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

func NewUser(name, email string) User{
	return User{
		Username: name,
		Email: email,
	}
}

func main() {
	user := User{}
	produit := Produit{}

	orm := orm.NewORM()
	orm.InitDB("mydb.db")
	orm.AutoMigrate(user, produit)

	//u := NewUser("Mouhamed Sylla","syllamouhamed99@gmail.com",)
	//u1 := NewUser("Abdou","abdou@gmail.com")
	//u2 := NewUser("Sidi","sidi@gmail.com")
	// // p := Produit{
	// // 	Name_produit: "Macbook Pro",
	// // 	Prix: 500000,
	// // }

	//orm.Insert(u1, u2)

	orm.SetModel("Email", "abdou@gmail.com", "User").UpdateField("moussa@gmail.com").Update(orm.Db)
}
