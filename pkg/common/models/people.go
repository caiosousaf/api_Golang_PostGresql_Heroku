package models


type Pessoa struct {
    ID_Pessoa       uint   `gorm:"primary_key" json:"id_pessoa"`
    Nome_Pessoa		string `gorm:"type: varchar(30) not null" json:"nome_pessoa"`
	Funcao_Pessoa	string `gorm:"type: varchar(15) not null" json:"funcao_pessoa"`
	EquipeID		int    `json:"equipeId"`
	Equipe			Equipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}