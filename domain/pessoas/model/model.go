package pessoas

import "time"

type ReqPessoa struct {
	ID_Pessoa        *int       `json:"id_pessoa" example:"4"`
	Nome_Pessoa      *string    `json:"nome_pessoa" example:"Caio Sousa"`
	Funcao_Pessoa    *string    `json:"funcao_pessoa" example:"Back-End"`
	Equipe_ID        *int       `json:"equipe_id" example:"1"`
	Data_Contratacao *time.Time `json:"data_contratacao" example:"19/09/2022"`
}

type ReqAtualizarPessoa struct {
	Nome_Pessoa   *string `json:"nome_pessoa" example:"Caio Swagger"`
	Funcao_Pessoa *string `json:"funcao_pessoa" example:"Back-End"`
	Equipe_ID     *int    `json:"equipe_id" example:"1"`
}

type ReqMembros struct {
	ID_Pessoa        *int       `json:"id_pessoa"`
	Nome_Pessoa      *string    `json:"nome_pessoa"`
	Funcao_Pessoa    *string    `json:"funcao_pessoa"`
	Equipe_ID        *int       `json:"equipe_id"`
	Data_Contratacao *time.Time `json:"data_contratacao"`
}

type ReqGetPessoa struct {
	ID_Pessoa        *uint   `json:"id_pessoa" example:"4"`
	Nome_Pessoa      *string `json:"nome_pessoa" example:"Caio Sousa"`
	Funcao_Pessoa    *string `json:"funcao_pessoa" example:"Back-End"`
	EquipeID         *int    `json:"equipe_id" example:"4"`
	Nome_Equipe      *string `json:"nome_equipe" example:"Komanda"`
	Data_Contratacao *string `json:"data_contratacao" example:"19/09/2022"`
}

type ReqTarefaPessoa struct {
	ID_Pessoa      *int       `json:"id_pessoa" example:"4"`
	Nome_Pessoa    *string    `json:"nome_pessoa" example:"Caio Sousa"`
	Funcao_Pessoa  *string    `json:"funcao_pessoa" example:"Back-End"`
	ID_Equipe      *int       `json:"id_equipe" example:"1"`
	Nome_Equipe    *string    `json:"nome_equipe" example:"Komanda"`
	Nome_Projeto   *string    `json:"nome_projeto" example:"Casas Bahias"`
	ID_Task        *int       `json:"id_task" example:"4"`
	Descricao_Task *string    `json:"descricao_task" example:"Exemplo"`
	Projeto_ID     *int       `json:"projeto_id" example:"40"`
	Status         *string    `json:"status" example:"Em Teste"`
	Data_Criacao   *time.Time `json:"data_criacao" example:"22/09/2022"`
	Data_Conclusao *time.Time `json:"data_conclusao" example:"22/09/2022"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega" example:"22/09/2022"`
	Prioridade     *int       `json:"prioridade" example:"2"`
}
