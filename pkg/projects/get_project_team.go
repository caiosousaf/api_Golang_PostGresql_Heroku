package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProjectTeam struct {
	Nome_Equipe 	string  `json:"nome_equipe"`
	ID_Projeto 		uint 	`json:"id_projeto"`
	Nome_Projeto 	string  `json:"nome_projeto"`
}

func (h handler) GetProjectTeam (c *gin.Context) {
	var equipes []GetProjectTeam

	if equipes := h.DB.Raw("select eq.nome_equipe, pr.id_projeto, pr.nome_projeto from equipes as eq inner join projetos as pr on eq.id_equipe = pr.equipe_id").Scan(&equipes); equipes.Error != nil {
		c.AbortWithError(http.StatusNotFound, equipes.Error)
		return
	}

	c.JSON(http.StatusOK, &equipes)
}