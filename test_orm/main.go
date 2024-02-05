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
	//user := User{}
	//produit := Produit{}

	orm := orm.NewORM()
	orm.InitDB("test.db")
	//orm.AutoMigrate(User{}, Produit{})

	// var produits []interface{}
	// for i := 1; i <= 100; i++ {
	// 	name := fmt.Sprintf("Produit%d", i)
	// 	produit := NewProduit(name, int64(i*5))
	// 	produits = append(produits, produit)
	// }

	// orm.Insert(produits...)

	orm.SetModel("Email", "syllamouhamed99@gmail.com", User{}).UpdateField("ahmed", "Username").Update(orm.Db)

	// orm.Delete(User{}, "Id", 2)
	//u := NewUser("Mouhamed Sylla", "syllamouhamed99@gmail.com")
	// u1 := NewUser("Abdou", "abdou@gmail.com")
	// u2 := NewUser("Sidi", "sidi@gmail.com")

	//orm.Insert(u)

	// orm.Custom.Where("Id", 2)
	// p := orm.Scan(Produit{}, "Prix", "Name_produit").([]struct {
	// 	Prix int64
	// 	Name_produit string
	// })

	// for _, v := range p {
	// 	fmt.Println(v.Prix)
	// 	fmt.Println(v.Name_produit)
	// }
}
