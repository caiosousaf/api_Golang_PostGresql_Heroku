package equipes

type ReqEquipe struct {
	ID_Equipe	*uint	`json:"id_equipe"`
	Nome_Equipe *string `json:"nome_equipe,omitempty"`
}