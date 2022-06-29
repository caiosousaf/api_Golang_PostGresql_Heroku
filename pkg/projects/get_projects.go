package projetos

import (
	"net/http"

	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
	"github.com/gin-gonic/gin"
)
type Projeto struct {
	ID_Projeto 		uint 	`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 	string 	`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID 		int 	`json:"equipeId"`
	Status			int		`json:"status"`
}

func (h handler) GetProjects(c *gin.Context) {
	var projetos []Projeto

	if result := h.DB.Find(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	c.JSON(http.StatusOK, &projetos)
}
