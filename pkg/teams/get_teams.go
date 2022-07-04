package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTeams(c *gin.Context) {
	var equipes []models.Equipe

	if result := h.DB.Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &equipes)
	
}
