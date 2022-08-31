package equipes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdateEquipeRequestBody struct {
	Nome_Equipe string `json:"nome_equipe"`
}

func (h handler) UpdateEquipe(c *gin.Context) {
	id := c.Param("id")
	body := UpdateEquipeRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	equipe.Nome_Equipe = body.Nome_Equipe

	h.DB.Save(&equipe)

	c.JSON(http.StatusOK, &equipe)
}