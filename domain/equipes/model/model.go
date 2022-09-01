package equipes

import "time"

type ReqEquipe struct {
	ID_Equipe	*uint	`json:"id_equipe"`
	Nome_Equipe *string `json:"nome_equipe,omitempty"`
	Data_Criacao *time.Time `json:"data_criacao"`
}

type ReqEquipeMembros struct {
	ID_Pessoa     *int    `json:"id_pessoa"`
	Nome_Pessoa   *string `json:"nome_pessoa"`
	Funcao_Pessoa *string `json:"funcao_pessoa"`
	Equipe_ID	  *int	  `json:"equipe_id"`
	Data_Contratacao	*time.Time `json:"data_contratacao"`
}

type ReqEquipeProjetos struct {
	Nome_Equipe  		string 		`json:"nome_equipe"`
	ID_Projeto   		uint   		`json:"id_projeto"`
	Nome_Projeto 		string 		`json:"nome_projeto"`
	Descricao_Projeto	string		`json:"descricao_projeto"`
	Status            	string     	`json:"status"`
    Data_Criacao      *time.Time     `json:"data_criacao"`
    Data_Conclusao    *time.Time    `json:"data_conclusao"`
    Prazo_Entrega     *time.Time 	`json:"prazo_entrega"`
}