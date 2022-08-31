package models

type Task struct {
	ID_Task         uint 	`gorm:"primary_key" json:"id_task"`
	Descricao_Task  string  `json:"descricao_task"`
	Nivel			string 	`json:"nivel"`
	PessoaID  		int		`json:"pessoa_id"`
	ProjetoID 		int 	`json:"projeto_id"`
	Status			string	`json:"status"`
}