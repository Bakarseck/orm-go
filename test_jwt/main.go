package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Bakarseck/orm-go/jwt"
	"github.com/Bakarseck/orm-go/utils"
)

type MyPayload struct {
	Name string
	Id   int
}

func (m *MyPayload) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func main() {
	utils.LoadEnv(".env")

	SECRET := os.Getenv("SECRET_KEY")

	header := jwt.Header{
		Alg: "alg",
		Typ: "jwt",
	}

	payload := &MyPayload{
		Name: "Bakar",
		Id:   12,
	}

	token, err := jwt.GenerateJWT(header, payload, SECRET)

	if err != nil {
		fmt.Println("error generating")
	}

	fmt.Println(token)

}
