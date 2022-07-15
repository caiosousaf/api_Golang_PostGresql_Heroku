package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID_Usuario		uint			`json:"id_usuario" gorm:"primaryKey"`
	Nome			string			`json:"nome"`
	Email			string			`json:"email"`
	Password		string			`json:"password"`
	CreatedAt		time.Time		`json:"created"`
	UpdateAt		time.Time		`json:"updated"`
	DeletedAt		gorm.DeletedAt	`gorm:"index"  json:"deleted"`
}