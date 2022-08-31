package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var host string = "ec2-3-223-169-166.compute-1.amazonaws.com"
var user string = "jikgljzwijqcwo"
var password string = "ba75e4994ed05c2580b76e88b34c09f8358b07b5dfcf4ee60ad58e37718c099d"
var dbname string = "d85bnr26iqqt4h"
const port = 5432

func Conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}