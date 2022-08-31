package pessoas

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPessoas(c *gin.Context) {
	var pessoas []models.Pessoa //Lista que vai ser usada para consultar as pessoas com os campos definidos no struct de Pessoa

	// Select da lista geral de pessoas
	if result := h.DB.Order("id_pessoa").Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Mostrando lista preenchida
	c.JSON(http.StatusOK, &pessoas)
}
