package pessoas

import "time"

type ReqPessoa struct {
	Nome_Pessoa      *string `json:"nome_pessoa"`
	Funcao_Pessoa    *string `json:"funcao_pessoa"`
	Equipe_ID        *int    `json:"equipe_id" `
	Data_Contratacao *time.Time `json:"data_contratacao"`
}