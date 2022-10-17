package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var host string = os.Getenv("DB_HOST")
var user string = os.Getenv("DB_USER")
var password string = os.Getenv("DB_PASSWORD")
var dbname string = os.Getenv("DB_NAME")
const port = 5432


func Conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}