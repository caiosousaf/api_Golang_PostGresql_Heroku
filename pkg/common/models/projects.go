package models

type Projeto struct {
	ID_Projeto 		uint 	`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 	string 	`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID 		int 	`json:"equipe_id"`
}