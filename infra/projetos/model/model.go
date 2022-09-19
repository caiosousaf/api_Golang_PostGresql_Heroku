package projetos


type ReqProjeto struct {
	Nome_Projeto      *string     `json:"nome_projeto"`
	Equipe_ID         *int        `json:"equipe_id"`
	Descricao_Projeto *string     `json:"descricao_projeto"`
	Prazo             int 		`json:"prazo_entrega"`
}

type ReqAtualizarProjetoData struct {
	Nome_Projeto      *string `json:"nome_projeto" example:"Casas Bahias"`
	Equipe_ID         *int    `json:"equipe_id" example:"1"`
	Descricao_Projeto *string `json:"descricao_projeto" example:"Criacao de sistema e-commerce"`
}

type ReqUpdateStatusProjeto struct {
	Status		*string		`json:"status" example:"Em Andamento"`
}