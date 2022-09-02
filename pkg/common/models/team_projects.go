package models

import "time"

type TeamProjects struct {
    ID_Projeto        uint   `gorm:"primary_key" json:"id_projeto"`
    Nome_Projeto      string `gorm:"type: varchar(30) not null" json:"nome_projeto"`
    EquipeID          int    `json:"equipe_id"`
    Status            string `json:"status"`
    Descricao_Projeto string `json:"descricao_projeto"`
	Prazo_Entrega	  *time.Time `json:"prazo_entrega"`
}