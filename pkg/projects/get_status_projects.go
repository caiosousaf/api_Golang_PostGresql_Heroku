package projetos

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get Projects with specific status
// @Summary Get Status of Projects with a specific status with Param Status
// @Description GET all registered projects that have the status passed as a parameter
// @Param        status   path      string  true  "Status"		Enums(A Fazer,Em Andamento,Em Teste,Concluido)
// @Accept json
// @Produce json
// @Success 200 {array} models.Projeto
// @Failure 400,404 {string} string "error"
// @Tags Projects
// @Router /projetos/status/{status} [get]
func (h handler) GetStatusProjects(c *gin.Context) {

	status := c.Param("status")

	var projeto []models.Projeto

	if result := h.DB.Raw("select * from projetos where status = ?", status).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)

}