package equipes

import (
	"net/http"
	"time"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	
)

type AddEquipeRequestBody struct {
	Nome_Equipe		string 			`json:"nome_equipe"`
}

func (h handler) AddTeam(c *gin.Context) {
	body := AddEquipeRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe models.Equipe
	dt := time.Now()

	equipe.Nome_Equipe = body.Nome_Equipe
	equipe.Data_Criacao = dt.Format("02-01-2006")

	if result := h.DB.Create(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &equipe)
}