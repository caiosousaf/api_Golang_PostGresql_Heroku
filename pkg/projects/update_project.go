package projetos

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

type UpdateProjetoRequestBody struct {
	Nome_Projeto	string `json:"nome_projeto"`
	Equipe 			models.Equipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}

func (h handler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	body := UpdateProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.Equipe = body.Equipe

	h.DB.Save(&projeto)

	c.JSON(http.StatusOK, &projeto)
}