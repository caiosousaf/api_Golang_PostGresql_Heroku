package projetos

import (
	"net/http"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetStatusProjects(c *gin.Context) {

	status := c.Param("status")

	var projeto []models.Projeto

	if result := h.DB.Raw("select * from projetos where status = ?", status).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)

}