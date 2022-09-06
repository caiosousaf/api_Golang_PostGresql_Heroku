package projetos


type ReqProjeto struct {
	Nome_Projeto      *string     `json:"nome_projeto"`
	Equipe_ID         *int        `json:"equipe_id"`
	Descricao_Projeto *string     `json:"descricao_projeto"`
	Prazo             int 		`json:"prazo_entrega"`
}