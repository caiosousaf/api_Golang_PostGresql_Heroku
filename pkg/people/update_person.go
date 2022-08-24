package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdatePessoaRequestBody struct {
	ID_Pessoa     uint   `json:"id_pessoa"`
	Nome_Pessoa   string `json:"nome_pessoa"`
	Funcao_Pessoa string `json:"funcao_pessoa"`
	Equipe_ID      int    `json:"equipe_id"`
}


// @Security bearerAuth
// @Summary PUT Person with ID
// @Description PUT a specific person. For the request to be met, the "nome_pessoa" and "funcao_pessoa" and "equipe_id" are required
// @Param        id   				path      	int  	true  	"Pessoa ID"
// @Param		Pessoa				body		string 	true 	"Pessoa"
// @Accept json
// @Produce json
// @Success 200 {object} UpdatePessoaRequestBody
// @Failure 400	{array} models.Error400Update
// @Failure 404 {array} models.Error404Update
// @Tags People
// @Router /pessoas/{id} [put]
func (h handler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "could not be modified. The parameters do not meet the requirements " + err.Error() ,
		})
		return
	}

	var pessoa models.Pessoa

	// Find first person with specific id
	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.EquipeID = body.Equipe_ID

	if result := h.DB.Raw("update pessoas set nome_pessoa = ?, funcao_pessoa = ?, equipe_id = ? where id_pessoa = ?", pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.EquipeID, pessoa.ID_Pessoa).Scan(&pessoa); result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Loss of contact with the database " ,
		})
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}