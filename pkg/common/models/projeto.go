package models

import "time"

type Projeto struct {
	ID_Projeto 			uint 	`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 		string 	`json:"nome_projeto"`
	Descricao_Projeto	string	`json:"descricao_projeto"`
	EquipeID 			int 	`json:"equipe_id"`
	Status				string	`json:"status"`
	DataInicio 			*time.Time 		`json:"data_inicio"`
	DataConclusao 		*time.Time		`json:"data_conclusao"`
}