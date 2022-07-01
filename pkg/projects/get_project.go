package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

type Projetos struct {
	ID_Projeto 			uint 		`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 		string 		`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID 			uint 		`json:"id_equipe"`
	Status				string		`json:"status"`
	Descricao_Projeto	string		`json:"descricao_projeto"`

}

func (h handler) GetProject(c *gin.Context) {
	id := c.Param("id")

	var projeto []Projetos

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)
}