package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID_Usuario uint           `json:"id_usuario" gorm:"primaryKey"`
	Nome       string         `json:"nome"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	CreatedAt  time.Time      `json:"created"`
	UpdateAt   time.Time      `json:"updated"`
	DeletedAt  gorm.DeletedAt `gorm:"index"  json:"deleted"`
}

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-23-23-151-191.compute-1.amazonaws.com"
	port     = 5432
	user     = "icsebrcphzbchf"
	password = "02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473"
	dbname   = "dcqvoffgfp6u50"
)

// cria e retorna uma conexão com o bando de dados postgres
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