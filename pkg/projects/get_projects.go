package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)


func (h handler) GetProjects(c *gin.Context) {
	var projetos []models.Projeto

	if result := h.DB.Raw("select * from projetos order by id_projeto").Scan(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	c.JSON(http.StatusOK, &projetos)
}
