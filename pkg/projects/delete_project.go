package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a specific Project
// @Description DELETE a specific project. For the request to be met, the "id_projeto" are required
// @Param        ID   						path      	int  	true  	"Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Projeto
// @Failure 400,404 {string} string "error"
// @Tags Projects
// @Router /projetos/{id} [delete]
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