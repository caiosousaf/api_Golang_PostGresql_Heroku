package models


type Pessoa struct {
	ID_Pessoa       uint   `json:"id_pessoa"`
    Nome_Pessoa		string `json:"nome_pessoa"`
	Funcao_Pessoa	string `json:"funcao_pessoa"`
	EquipeID		int    `json:"equipeId"`
}