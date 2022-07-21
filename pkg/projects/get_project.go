package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"

)

func (h handler) GetProject(c *gin.Context) {
	id := c.Param("id")

	var projeto []Projects

	if result := h.DB.Raw("select * from projetos where id_projeto = ?", id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)
}