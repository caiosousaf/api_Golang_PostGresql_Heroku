package projetos

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

func (h handler) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	var projeto models.Projeto

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&projeto)

	c.Status(http.StatusOK)
}