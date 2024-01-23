package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

type Model struct {
	ID        int64      `orm-go:"PRIMARY KEY AUTOINCREMENT"`
	CreatedAt time.Time `orm-go:"DEFAULT CURRENT_TIMESTAMP"`
}

type ORM struct {
	db *sql.DB
}

func NewORM() *ORM {
	return &ORM{}
}

func (o *ORM) InitDB(name string) {
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

	sqlTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", v.Name())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		fieldType := v.Field(i).Type
		if fieldType.Kind() == reflect.Struct {
			for i := 0; i < fieldType.NumField(); i++ {
				struct_field := fieldType.Field(i)
				ormgoTag := struct_field.Tag.Get("orm-go")
				sqlTable += "\t" + struct_field.Name + " " + GetType(struct_field.Type) + " " + ormgoTag + ",\n"
			}
		} else {
			ormgoTag := field.Tag.Get("orm-go")

			if i == v.NumField()-1 {
				sqlTable += "\t" + field.Name + " " + GetType(fieldType) + " " + ormgoTag + "\n"
			} else {
				sqlTable += "\t" + field.Name + " " + GetType(fieldType) + " " + ormgoTag + ",\n"
			}
		}

		
	}
	sqlTable += ")"
	fmt.Println(sqlTable)

	// _, err := o.db.Exec(sqlTable)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
