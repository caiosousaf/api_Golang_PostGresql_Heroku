package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type GetMembers struct {
	ID_Equipe   uint             `json:"id_equipe"`
	Nome_Equipe string           `json:"nome_equipe"`
	Pessoas     []models.Pessoa  `json:"pessoas"`
	Projetos    []models.Projeto `json:"projetos"`
}

func (h handler) GetTeams(c *gin.Context) {
	var equipes []models.Equipe
	var eq []GetMembers

	if result := h.DB.Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := 0; i < len(equipes); i++ {
		var pessoas []models.Pessoa
		var projetos []models.Projeto
		if result := h.DB.Raw("select * from pessoas where equipe_id = ?", equipes[i].ID_Equipe).Scan(&pessoas); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		if result := h.DB.Raw("select * from projetos where equipe_id = ?", equipes[i].ID_Equipe).Scan(&projetos); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
		}

		e := &GetMembers{
			ID_Equipe:   equipes[i].ID_Equipe,
			Nome_Equipe: equipes[i].Nome_Equipe,
			Pessoas:     pessoas,
			Projetos:    projetos,
		}
		eq = append(eq, *e)
	}
	c.JSON(http.StatusOK, &eq)

}
