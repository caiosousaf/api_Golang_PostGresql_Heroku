package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary DELETE a Team
// @Description DELETE a Team
// @Param		id		path	int		true		"Team_ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Equipe
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /equipes/{id} [delete]
func (h handler) DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&equipe)

	c.Status(http.StatusOK)
}