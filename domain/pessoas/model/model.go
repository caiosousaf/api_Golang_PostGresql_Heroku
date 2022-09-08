package pessoas

import "time"

type ReqPessoa struct {
	ID_Pessoa        *int       `json:"id_pessoa"`
	Nome_Pessoa      *string    `json:"nome_pessoa"`
	Funcao_Pessoa    *string    `json:"funcao_pessoa"`
	Equipe_ID        *int       `json:"equipe_id" `
	Data_Contratacao *time.Time `json:"data_contratacao"`
}

type ReqAtualizarPessoa struct {
	Nome_Pessoa   *string `json:"nome_pessoa"`
	Funcao_Pessoa *string `json:"funcao_pessoa"`
	Equipe_ID     *int    `json:"equipe_id" `
}

type ReqMembros struct {
	ID_Pessoa        *int       `json:"id_pessoa"`
	Nome_Pessoa      *string    `json:"nome_pessoa"`
	Funcao_Pessoa    *string    `json:"funcao_pessoa"`
	Equipe_ID        *int       `json:"equipe_id"`
	Data_Contratacao *time.Time `json:"data_contratacao"`
}

type ReqGetPessoa struct {
	ID_Pessoa        *uint   `json:"id_pessoa"`
	Nome_Pessoa      *string `json:"nome_pessoa"`
	Funcao_Pessoa    *string `json:"funcao_pessoa"`
	EquipeID         *int    `json:"equipe_id"`
	Nome_Equipe      *string `json:"nome_equipe"`
	Data_Contratacao *string `json:"data_contratacao"`
}

type ReqTarefaPessoa struct {
	ID_Pessoa      *int       `json:"id_pessoa"`
	Nome_Pessoa    *string    `json:"nome_pessoa"`
	Funcao_Pessoa  *string    `json:"funcao_pessoa"`
	ID_Equipe      *int       `json:"id_equipe"`
	Nome_Equipe    *string    `json:"nome_equipe"`
	Nome_Projeto   *string    `json:"nome_projeto"`
	ID_Task        *int       `json:"id_task"`
	Descricao_Task *string    `json:"descricao_task"`
	Projeto_ID     *int       `json:"projeto_id"`
	Status         *string    `json:"status"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     *int       `json:"prioridade"`
}
