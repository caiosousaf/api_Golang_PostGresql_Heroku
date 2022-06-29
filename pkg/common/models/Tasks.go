package models

type Task struct {
	ID_Task         uint `gorm:"primary_key" json:"id_task"`
	Descricao_Task  string `json:"descricao_task"`
	PessoaID  		int		`json:"id_pessoa"`
	ProjetoID 		int 	`json:"id_projeto"`
}