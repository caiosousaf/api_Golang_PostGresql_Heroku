package database

import (
	"database/sql"
	"fmt"
	//"os"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-23-23-151-191.compute-1.amazonaws.com"
	port     = 5432
	user     = "icsebrcphzbchf"
	password = "02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473"
	dbname   = "dcqvoffgfp6u50"
)

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