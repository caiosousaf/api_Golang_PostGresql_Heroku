package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Equipe struct {
    ID_Equipe         uint   `gorm:"primary_key" json:"id_equipe"`
    Nome_Equipe       string `json:"nome_equipe"`
	Membros			models.Pessoa
}

func (h handler) GetTeams(c *gin.Context) {
	var equipes []Equipe

	if result := h.DB.Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &equipes)
	
}
