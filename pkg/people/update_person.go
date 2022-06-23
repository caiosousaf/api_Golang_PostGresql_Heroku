package pessoas

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

type UpdatePessoaRequestBody struct {
	Nome_Pessoa		string `json:"nome_pessoa"`
	Funcao_Pessoa 		string `json:"funcao_pessoa"`
	Equipe 		models.Equipe `json:"equipe"`
}

func (h handler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.Equipe = body.Equipe

	h.DB.Save(&pessoa)

	c.JSON(http.StatusOK, &pessoa)
}