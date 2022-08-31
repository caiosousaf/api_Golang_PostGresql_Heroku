package models

import "time"

type Pessoa struct {
    ID_Pessoa       uint			`gorm:"primary_key" json:"id_pessoa"`
    Nome_Pessoa		string			`json:"nome_pessoa"`
	Funcao_Pessoa	string			`json:"funcao_pessoa"`
	EquipeID		int				`json:"equipe_id"`
	Favoritar		uint			`json:"favoritar"`
	DataContratacao time.Time		`json:"data_contratacao"`
}