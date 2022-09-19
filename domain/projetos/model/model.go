package projetos

import "time"

type ReqProjeto struct {
	Nome_Projeto      *string `json:"nome_projeto" example:"Nome"`
	Equipe_ID         *int    `json:"equipe_id" example:"2"`
	Descricao_Projeto *string `json:"descricao_projeto" example:"Descricao"`
	Prazo             int     `json:"prazo_entrega" example:"2"`
}

type ReqProjetos struct {
	ID_Projeto        *uint      `json:"id_projeto" example:"58"`
	Nome_Projeto      *string    `json:"nome_projeto" example:"Nome"`
	Descricao_Projeto *string    `json:"descricao_projeto" example:"Descricao"`
	EquipeID          *int       `json:"equipe_id" example:"2"`
	Nome_Equipe       *string    `json:"nome_equipe" example:"Cariri Inovação"`
	Status            *string    `json:"status" example:"Concluido"`
	Data_Criacao      *string    `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    *string    `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}

type ReqStatusProjeto struct {
	ID_Projeto        *uint      `json:"id_projeto"`
	Nome_Projeto      *string    `json:"nome_projeto"`
	EquipeID          *int       `json:"equipe_id"`
	Status            *string    `json:"status"`
	Descricao_Projeto *string    `json:"descricao_projeto"`
	Data_Criacao      *time.Time `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    *time.Time `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}

type ReqAtualizarProjeto struct {
	ID_Projeto        *uint      `json:"id_projeto" example:"58"`
	Nome_Projeto      *string    `json:"nome_projeto" example:"Nome"`
	Descricao_Projeto *string    `json:"descricao_projeto" example:"Descricao"`
	EquipeID          *int       `json:"equipe_id" example:"2"`
	Status            *string    `json:"status" example:"Concluido"`
	Data_Criacao      *string    `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    *string    `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}

type ReqTasksProjeto struct {
	ID_Projeto     *uint      `json:"id_projeto"`
	Nome_Projeto   *string    `json:"nome_projeto"`
	Nome_Equipe    *string    `json:"nome_equipe"`
	ID_Task        *int       `json:"id_task"`
	Descricao_Task *string    `json:"descricao_task"`
	Projeto_ID     *int       `json:"projeto_id"`
	Pessoa_ID      *int       `json:"pessoa_id"`
	Nome_Pessoa    *string    `json:"nome_pessoa"`
	Status         *string    `json:"status"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     *int       `json:"prioridade"`
}
