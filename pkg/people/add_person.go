package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddPessoaRequestBody struct {
	Nome_Pessoa   string `json:"nome_pessoa" example:"Caio Sousa"`
	Funcao_Pessoa string `json:"funcao_pessoa" example:"Back-End Developer"`
	Equipe_ID     int    `json:"equipe_id" example:"1"`
}

// @Security bearerAuth
// @Summary POST a new Person
// @Description POST a new Person. For the request to be met, the "nome_pessoa", "funcao_pessoa", "equipe_id" are required. 
// @Param		Person		body	string		true	"New Person"
// @Accept json
// @Produce json
// @Success 201 {object} AddPessoaRequestBody
// @Failure 400 {array} models.Error400Create
// @Failure 404 {array} models.Error404Create
// @Tags People
// @Router /pessoas [post]
func (h handler) AddPerson(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error() ,
		})
		return
	}
	
	var pessoa models.Pessoa

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.EquipeID = body.Equipe_ID

	// Function that creates a new person
	if result := h.DB.Create(&pessoa).Scan(&pessoa); result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Loss of contact with the database " ,
		})
		return
	}

	// Show what you just created
	c.JSON(http.StatusCreated, &pessoa)
}