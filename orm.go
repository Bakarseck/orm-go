package orm

import "database/sql"

type ORM struct {
	//db *sql.DB
}

func (o *ORM) InitDB(name string) (db *sql.DB){
	
	return
}


func (o *ORM) AutoMigrate(table interface{}) {
	
}