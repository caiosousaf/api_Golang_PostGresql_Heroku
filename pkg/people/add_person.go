package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddPessoaRequestBody struct {
	Nome_Pessoa   string `json:"nome_pessoa"`
	Funcao_Pessoa string `json:"funcao_pessoa"`
	Equipe_ID     int    `json:"equipe_id"`
}

// @Summary POST a new Person
// @Description POST a new Person. For the request to be met, the "nome_pessoa", "funcao_pessoa", "equipe_id" are required. 
// @Param		Person		body	string		true	"New Person"
// @Accept json
// @Produce json
// @Success 200 {object} AddPessoaRequestBody
// @Failure 400,404 {string} string "error"
// @Tags People
// @Router /pessoas [post]
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
	pessoa.EquipeID = body.Equipe_ID

	if result := h.DB.Create(&pessoa).Scan(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pessoa)
}