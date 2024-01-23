package main

import (
	"github.com/Bakarseck/orm-go"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	orm.Model
	Username string `orm-go:"not null;unique"`
	Email    string `orm-go:"not null;unique"`
}

func main() {
	// var user User
	orm := orm.NewORM()
	orm.InitDB("test.db")
	// orm.AutoMigrate(user)
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/Bakarseck/orm-go/jwt"
// 	"github.com/Bakarseck/orm-go/utils"
// )

// type MyPayload struct {
// 	Sub  string `json:"sub"`
// 	Name string `json:"name"`
// 	Iat  int64  `json:"iat"`
// }

// func (p *MyPayload) ToJSON() ([]byte, error) {
// 	return json.Marshal(p)
// }

// func main() {
// 	utils.LoadEnv(".env")

// 	SECRET := os.Getenv("SECRET")

// 	header := jwt.Header{
// 		Alg: "HS256",
// 		Typ: "JWT",
// 	}

//     payload := &MyPayload{
//         Sub:  "1234567890",
//         Name: "Bakar SECK",
//         Iat:  1516239022,
//     }

// 	jwt, err := jwt.GenerateJWT(header, payload, SECRET)

// 	if err != nil {
// 		log.Println("generate token failed:", err.Error())
// 	}

// 	fmt.Println(jwt)

// }
