package tasks

type ReqTaskData struct {
	Descricao_Task *string `json:"descricao_task" example:"Descrição Teste"`
	PessoaID       *int    `json:"pessoa_id" example:"4"`
	ProjetoID      *int    `json:"projeto_id" example:"24"`
	Prazo          int     `json:"prazo_entrega" example:"17"`
	Prioridade     *int    `json:"prioridade" example:"1"`
}
