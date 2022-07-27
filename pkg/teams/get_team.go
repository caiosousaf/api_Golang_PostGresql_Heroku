package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

func (h handler) GetTeam(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe
	var me GetMembers

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var pessoas []models.Pessoa
	if result := h.DB.Where("equipe_id = ?", equipe.ID_Equipe).Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	me.ID_Equipe = equipe.ID_Equipe
	me.Nome_Equipe = equipe.Nome_Equipe
	me.Pessoas = pessoas

	c.JSON(http.StatusOK, &me)
}