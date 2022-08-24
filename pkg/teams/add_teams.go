package equipes

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddEquipeRequestBody struct {
	Nome_Equipe string `json:"nome_equipe" example:"Krutaya Komanda"`
}

// @Security bearerAuth
// @Summary POST a new Team
// @Description POST a new Team. For the request to be met, the "nome_equipe" are required. 
// @Param		Team		body	string		true	"NewTeam"
// @Accept json
// @Produce json
// @Success 200 {object} AddEquipeRequestBody
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /equipes [post]
func (h handler) AddTeam(c *gin.Context) {
	body := AddEquipeRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe models.Equipe

	equipe.Nome_Equipe = body.Nome_Equipe

	if result := h.DB.Create(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &equipe)
}