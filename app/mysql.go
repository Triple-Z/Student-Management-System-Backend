package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func db_init() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8")
	checkError(err)

	err = db.Ping()
	checkError(err)

	log.Println("DB connect successfully")

	db.Exec("create database if not exists pdbs_course character set utf8")
	db.Exec("use pdbs_course")

	log.Println("DB database connected")

	db.Exec("create table if not exists department (id int auto_increment primary key, name varchar(50) unique, create_date datetime default current_timestamp, last_updated datetime default current_timestamp on update current_timestamp) default charset=utf8")

	log.Println("DB department table all set up")

	db.Exec("create table if not exists student (id int auto_increment primary key, number char(9), name char(10), department int, create_date datetime default current_timestamp, last_updated datetime default current_timestamp on update current_timestamp, foreign key (department) references department(id) on update cascade on delete set null) default charset=utf8")

	log.Println("DB student table all set up")

	return db
}
