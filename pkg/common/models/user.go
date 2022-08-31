package models

import "time"


type User struct {
    ID_User		uint		`gorm:"primary_key" json:"id_user"`
    Nome_User	string		`json:"nome_user"`
	Email		string		`json:"email"`
	Senha		string		`json:"senha"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}