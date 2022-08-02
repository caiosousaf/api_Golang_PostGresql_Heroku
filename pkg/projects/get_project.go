package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

// Get Project
// @Summary Get Project with specific ID
// @Description GET a project with a specific ID
// @Param        id   path      int  true  "Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} Projects
// @Failure 400,404 {string} string "error"
// @Tags Projects
// @Router /projetos/{id} [get]
func (h handler) GetProject(c *gin.Context) {
	id := c.Param("id")

	var projeto []Projects

	if result := h.DB.Raw("select * from projetos where id_projeto = ?", id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)
}