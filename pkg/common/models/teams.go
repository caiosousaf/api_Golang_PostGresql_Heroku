package models

type Equipe struct {
	ID_Equipe   uint   `gorm:"primary_key" json:"id_equipe"`
	Nome_Equipe string `json:"nome_equipe"`
}