package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	Nome_Projeto	string 			`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	Equipe 			models.Equipe	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}

func (h handler) AddProject(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.Equipe = body.Equipe

	if result := h.DB.Create(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &projeto)
}