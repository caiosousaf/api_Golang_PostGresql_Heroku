package tasks

import "time"

type ReqTask struct {
	ID_Task        *uint      `json:"id_task"`
	Descricao_Task *string    `json:"descricao_task"`
	PessoaID       *int       `json:"pessoa_id"`
	ProjetoID      *int       `json:"projeto_id"`
	Status         *string    `gorm:"varchar(20)" json:"status"`
	Data_Criacao   *string    `json:"data_criacao"`
	Data_Conclusao *string    `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     *int       `json:"prioridade"`
}

type ReqTaskApresent struct {
	Descricao_Task *string `json:"descricao_task" example:"Descrição Teste"`
	PessoaID       *int    `json:"pessoa_id" example:"4"`
	ProjetoID      *int    `json:"projeto_id" example:"24"`
	Prazo          int     `json:"prazo_entrega" example:"17"`
	Prioridade     *int    `json:"prioridade" example:"1"`
}

type ReqTasks struct {
	ID_Task        *uint       `json:"id_task" example:"10"`
	Descricao_Task *string     `json:"descricao_task" example:"Descricao"`
	PessoaID       *int        `json:"pessoa_id" example:"4"`
	Nome_Pessoa    *string     `json:"nome_pessoa" example:"Fulano"`
	ProjetoID      *int        `json:"projeto_id" example:"3"`
	Nome_Projeto   *string     `json:"nome_projeto" example:"Grupo BrisaNet"`
	Status         *string     `json:"status" example:"Em Andamento"`
	Data_Criacao   *time.Time `json:"data_criacao" example:"19/09/2022"`
	Data_Conclusao *time.Time `json:"data_conclusao" example:"19/09/2022"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega" example:"19/09/2022"`
	Prioridade     *int        `json:"prioridade" example:"2"`
}


type ResTasksbyTeam struct {
    ID_Task        uint       `json:"id_task"`
    Descricao_Task string     `json:"descricao_task"`
    Status         string     `json:"status" enums:"Em Andamento, Concluido"`
    Pessoa_ID      uint       `json:"pessoa_id"`
    Nome_Pessoa    string     `json:"nome_pessoa"`
    Projeto_ID     uint       `json:"projeto_id"`
    Data_Criacao   *time.Time `json:"data_criacao"`
    Data_Conclusao *time.Time `json:"data_conclusao"`
    Prazo_Entrega  *time.Time `json:"prazo_entrega"`
    Prioridade     int        `json:"prioridade"`
}