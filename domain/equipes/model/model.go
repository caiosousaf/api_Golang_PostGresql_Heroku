package equipes

type ReqEquipe struct {
	ID_Equipe	*uint	`json:"id_equipe"`
	Nome_Equipe *string `json:"nome_equipe,omitempty"`
}

type ReqEquipeMembros struct {
	ID_Pessoa     *int    `json:"id_pessoa"`
	Nome_Equipe   *string `json:"nome_equipe"`
	ID_Equipe     *uint    `json:"id_equipe"`
	Nome_Pessoa   *string `json:"nome_pessoa"`
	Funcao_Pessoa *string `json:"funcao_pessoa"`
}