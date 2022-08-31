package equipes

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetEquipes(c *gin.Context) {
	var equipes []models.Equipe

	if result := h.DB.Order("id_equipe").Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &equipes)
}