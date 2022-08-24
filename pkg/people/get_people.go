package pessoas

import (
	"net/http"

	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// Get People
// @Summary Get All People
// @Description Returns all registered people and the name of the team they are assigned to
// @Accept json
// @Produce json
// @Success 200 {array} GetPessoa
// @Failure 404 {array} models.Error404Get
// @Tags People
// @Router /pessoas [get]
func (h handler) GetPeople(c *gin.Context) {
	var pessoas []GetPessoa

	if pessoas := h.DB.Raw(`select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao
	from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe order by pe.id_pessoa`).Scan(&pessoas); pessoas.Error != nil {
		c.JSON(404, gin.H{
			"message": "Loss of contact with the database" ,
		})
		return
	}
	c.JSON(http.StatusOK, &pessoas)
}
