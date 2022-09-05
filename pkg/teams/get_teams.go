package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)



type GetMembers struct {
	ID_Equipe   uint             			`json:"id_equipe"`
	Nome_Equipe string           			`json:"nome_equipe"`
	Pessoas     []models.Pessoa  			`json:"pessoas"`
	Projetos_Ativos	string						`json:"projetos_ativos"`		
}

// Get Teams
// @Security bearerAuth
// @Summary Get All Teams
// @Description Returns all registered teams and all their members they are assigned to
// @Accept json
// @Produce json
// @Success 200 {array} GetMembers
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /equipes [get]
func (h handler) GetTeams(c *gin.Context) {
	var equipes []models.Equipe
	var eq []GetMembers

	if result := h.DB.Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := 0; i < len(equipes); i++ {
		var pessoas []models.Pessoa
		if result := h.DB.Raw("select * from pessoas where equipe_id = ?", equipes[i].ID_Equipe).Scan(&pessoas); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		var check string
		if result := h.DB.Raw(`select count(*) as Projetos_Ativos 
		from projetos where status = 'Em Andamento' and equipe_id = ?`, equipes[i].ID_Equipe).Scan(&check); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		e := &GetMembers{
			ID_Equipe:   equipes[i].ID_Equipe,
			Nome_Equipe: equipes[i].Nome_Equipe,
			Pessoas:     pessoas,
			Projetos_Ativos: check,
		}
		eq = append(eq, *e)
	}
	c.JSON(http.StatusOK, &eq)

}
