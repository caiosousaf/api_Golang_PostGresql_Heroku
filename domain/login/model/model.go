package login

import "time"

type Login struct {
	ID_Usuario   uint      `json:"id_usuario"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Data_Criacao time.Time `json:"data_criacao"`
}
