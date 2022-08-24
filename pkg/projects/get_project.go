package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

// @Security bearerAuth
// Get Project
// @Summary Get Project with specific ID
// @Description GET a project with a specific ID
// @Param        id   path      int  true  "Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} Projects	"OK"
// @Failure 400 {array} models.Error400Get
// @Failure 404 {array} models.Error404Get
// @Tags Projects
// @Router /projetos/{id} [get]
func (h handler) GetProject(c *gin.Context) {
	id := c.Param("id")
	var IdExist int

	if result := h.DB.Raw("select count(*) from projetos where id_projeto = ?", id).Scan(&IdExist); result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Cannot BindJSON. Type of ID-Project Invalid"  ,
		})
		return
	}

	var projeto []Projects

	if IdExist == 1 {
			if result := h.DB.Raw(`select pr.*, eq.nome_equipe from projetos pr inner join equipes eq
	 		on pr.equipe_id = eq.id_equipe where id_projeto = ?`, id).Scan(&projeto); result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Loss of contact with the database " ,
		})
		return
		}

	c.JSON(http.StatusOK, &projeto)
	} else {
		c.JSON(400, gin.H{
			"message": "Data not found with the passed parameters " ,
		})
		return
	}


}