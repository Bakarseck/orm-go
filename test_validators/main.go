package main

import (
	"fmt"

	"github.com/Bakarseck/orm-go/validators"
)

type User struct {
	Username string `validate:"required,username"`
	Password string `validate:"required,password"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=18,max=99"`
}

func main() {

	user := User{
		Username: "my_username123",
		Password: "Passw0rd§",
		Email:    "user@example.com",
		Age:      18,
	}

	results := validators.ValidateStruct(user)

	for _, result := range results {
		if result.Valid {
			fmt.Printf("Champ '%s' avec tag '%s' est valide.\n", result.Field, result.Tag)
		} else {
			fmt.Printf("Champ '%s' avec tag '%s' n'est pas valide: %s\n", result.Field, result.Tag, result.Reason)
		}
	}
}
