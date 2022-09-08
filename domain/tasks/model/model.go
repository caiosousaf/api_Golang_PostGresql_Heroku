package tasks

import "time"

type ReqTask struct {
	ID_Task        *uint      `gorm:"primary_key" json:"id_task"`
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
