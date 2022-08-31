package equipes

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

func (h handler) DeleteEquipe(c *gin.Context) {
	id := c.Param("id") // Pegando o ID como parâmetro da URL

	var equipe models.Equipe // Declarando uma variável do tipo equipe que será usada para a consulta

	// Achando o registro de equipe referente ao ID informado na URL
	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Comando para deletar o registro especificado pelo ID
	h.DB.Delete(&equipe)

	c.Status(http.StatusOK)
}