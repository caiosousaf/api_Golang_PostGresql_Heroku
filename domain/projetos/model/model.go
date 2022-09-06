package projetos

import "time"

type ReqProjeto struct {
	Nome_Projeto      *string `json:"nome_projeto"`
	Equipe_ID         *int    `json:"equipe_id"`
	Descricao_Projeto *string `json:"descricao_projeto"`
	Prazo             int     `json:"prazo_entrega"`
}

type ReqProjetos struct {
	ID_Projeto        *uint      `gorm:"primary_key" json:"id_projeto" example:"58"`
	Nome_Projeto      *string    `gorm:"type: varchar(30) not null" json:"nome_projeto" example:"Nome"`
	Descricao_Projeto *string    `json:"descricao_projeto" example:"Descricao"`
	EquipeID          *int       `json:"equipe_id" example:"2"`
	Nome_Equipe       *string    `json:"nome_equipe" example:"Cariri Inovação"`
	Status            *string    `json:"status" example:"Concluido"`
	Data_Criacao      *string    `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    *string    `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}

type ReqStatusProjeto struct {
	ID_Projeto        *uint       `gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto      *string     `gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID          *int        `json:"equipe_id"`
	Status            *string     `json:"status"`
	Descricao_Projeto *string     `json:"descricao_projeto"`
	Data_Criacao      *time.Time `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    *time.Time `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}
