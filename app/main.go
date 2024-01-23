package main

<<<<<<< Updated upstream
import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Bakarseck/orm-go/jwt"
	"github.com/Bakarseck/orm-go/utils"
)

type MyPayload struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Iat  int64  `json:"iat"`
}

func (p *MyPayload) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func main() {
	utils.LoadEnv(".env")

	SECRET := os.Getenv("SECRET")

	header := jwt.Header{
		Alg: "HS256",
		Typ: "JWT",
	}

    payload := &MyPayload{
        Sub:  "1234567890",
        Name: "Bakar SECK",
        Iat:  1516239022,
    }

	jwt, err := jwt.GenerateJWT(header, payload, SECRET)

	if err != nil {
		log.Println("generate token failed:", err.Error())
	}

	fmt.Println(jwt)

=======
import "github.com/Bakarseck/orm-go"

func main() {
	//db := orm.NewORM().InitDB("mydb")

	orm.NewORM().AutoMigrate(orm.User{})
>>>>>>> Stashed changes
}
