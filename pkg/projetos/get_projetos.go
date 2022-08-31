package projetos

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjetos(c *gin.Context) {
	var projetos []models.Projeto

	if result := h.DB.Order("id_projeto").Find(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	c.JSON(http.StatusOK, &projetos)
}
