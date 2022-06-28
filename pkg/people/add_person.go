package pessoas

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

type AddPessoaRequestBody struct {
    Nome_Pessoa		string `json:"nome_pessoa"`
	Funcao_Pessoa	string `json:"funcao_pessoa"`
	EquipeID		int    `json:"equipeId"`
}

func (h handler) AddPerson(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.EquipeID = body.EquipeID

	if result := h.DB.Raw(`insert into pessoas(nome_pessoa, funcao_pessoa, equipe_id) values(?, ?, ?)`, pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.EquipeID).Scan(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pessoa)
}