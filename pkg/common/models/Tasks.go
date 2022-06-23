package models

type Task struct {
	ID_Task         uint `gorm:"primary_key" json:"id_task"`
	Descricao_Task  string `gorm:"type: varchar(100) not null" json:"descricao_task"`
	PessoaID  		int
	Pessoa			Pessoa `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"id_pessoa"`
	ProjetoID 		int 
	Projeto			Projeto `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"id_projeto"`
}