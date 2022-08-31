package equipes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type EquipesGetBody struct {
	ID_Equipe 	uint 				`json:"id_equipe"`
	Nome_Equipe string 				`json:"nome_equipe"`
	Pessoas 	[]models.Pessoa 	`json:"pessoas"`
}

func (h handler) GetEquipe(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe
	var eq EquipesGetBody

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var pessoas []models.Pessoa
	if result := h.DB.Where("equipe_id = ?", equipe.ID_Equipe).Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	eq.ID_Equipe = equipe.ID_Equipe
	eq.Nome_Equipe = equipe.Nome_Equipe
	eq.Pessoas = pessoas

	c.JSON(http.StatusOK, &eq)
}