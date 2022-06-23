package models

type Projeto struct {
	ID_Projeto 		uint 	`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 	string 	`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID 		int 	`json:"id_equipe"`
	Equipe 			Equipe 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}