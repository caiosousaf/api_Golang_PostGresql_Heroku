package models

type Task struct {
	ID_Task         uint `gorm:"primary_key" json:"id_task"`
	Descricao_Task  string `json:"descricao_task"`
	PessoaID  		int		`json:"pessoa_id"`
	ProjetoID 		int 	`json:"projeto_id"`
	Status			string		`gorm:"varchar(20)" json:"status"`
	Data_Criacao	string	`json:"data_criacao"`
}