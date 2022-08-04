package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateEquipeRequestBody struct {
	Nome_Equipe string `json:"nome_equipe"`
}

// PUT Project
// @Summary PUT Project with ID
// @Description PUT a specific project. For the request to be met, the "nome_projeto" and "equipe_id" and "descricao_projeto" are required
// @Param        id   				path      	int  	true  	"Team ID"
// @Param		Team				body		string 	true 	"Team"
// @Accept json
// @Produce json
// @Success 200 {object} UpdateEquipeRequestBody
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /teams/{id} [put]
func (h handler) UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	body := UpdateEquipeRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	equipe.Nome_Equipe = body.Nome_Equipe

	h.DB.Save(&equipe)

	c.JSON(http.StatusOK, &equipe)
}