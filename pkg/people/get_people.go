package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type GetPessoasRequestBody struct {
	ID_Pessoa 		uint	`json:"id_pessoa"`
    Nome_Pessoa		string 	`json:"nome_pessoa"`
	Funcao_Pessoa	string 	`json:"funcao_pessoa"`
	EquipeID		int 	`json:"equipe_id"`
	Data_Contratacao string	`json:"data_contratacao"`
}

func (h handler) GetPeople(c *gin.Context) {
	var pessoas []models.Pessoa

	if result := h.DB.Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var exibe []GetPessoasRequestBody
	p := GetPessoasRequestBody{}

	for _, pe := range pessoas{
		p.ID_Pessoa = pe.ID_Pessoa
		p.Nome_Pessoa = pe.Nome_Pessoa
		p.Funcao_Pessoa = pe.Funcao_Pessoa
		p.EquipeID = pe.EquipeID
		p.Data_Contratacao = pe.Data_Contratacao

		exibe = append(exibe, p)
	}


	c.JSON(http.StatusOK, &exibe)
}
