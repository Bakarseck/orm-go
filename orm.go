package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Model struct {
	ID        int  `orm-go:"primaryKey;autoIncrement"`
	CreatedAt Time `org-go:"current_timestamp"`
}



type ORM struct {
	db *sql.DB
}

func NewORM() *ORM {
	return &ORM{}
}

func (o *ORM) InitDB(name string) {
	fmt.Println("ddd")
	_, err := os.Stat(name)


	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	//database connection
	o.db, err = sql.Open("sqlite3", name)

	if err != nil {
		log.Fatal(err)
	}
}

func (o *ORM) AutoMigrate(table interface{}) {
	v := reflect.TypeOf(table)

	sqlTable := ""
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		fieldType := v.Field(i).Type
		gormTag := field.Tag.Get("orm-go")

		sqlTable += field.Name + " " + GetType(fieldType) + " " + gormTag + "\n"

	}

	fmt.Println(sqlTable)

	_, err := o.db.Exec(sqlTable)
	if err != nil {
		fmt.Println(err)
		return
	}
}
