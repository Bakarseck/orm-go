package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Model struct {
	ID       int   `orm-go:"primaryKey;autoIncrement"`
	CreatedAt Time `org-go:"current_timestamp"`
}


type User struct {
	Model
	Username string `orm-go:"not null;unique"`
	Email    string `orm-go:"not null;unique"`
}

type ORM struct {
	//db *sql.DB
}

func NewORM() *ORM {
	return &ORM{}
}

func (o *ORM) InitDB(name string) (db *sql.DB) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	//database connection
	db, err = sql.Open("sqlite3", name)

	if err != nil {
		log.Fatal(err)
	}
	return
}

func (o *ORM) AutoMigrate(table interface{}) {
	v := reflect.TypeOf(table)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		gormTag := field.Tag.Get("orm-go")
		fmt.Println(gormTag)
	}

}
