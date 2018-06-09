package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func db_init() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pdbs_course?parseTime=true&charset=utf8")
	checkError(err)

	err = db.Ping()
	checkError(err)

	log.Println("DB connect successfully")

	return db
}

