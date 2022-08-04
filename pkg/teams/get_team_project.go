package equipes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EquipeProjetos struct {
	Nome_Equipe  string `json:"nome_equipe"`
	ID_Projeto   uint   `json:"id_projeto"`
	Nome_Projeto string `json:"nome_projeto"`
}

// @Summary Get Projects of a specific Team
// @Description GET all projects of a specific Team with ID
// @Param        id   path      int  true  "Team ID"
// @Accept json
// @Produce json
// @Success 200 {array} EquipeProjetos
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /equipes/{id}/projetos [get]
func (h handler) GetTeamProject(c *gin.Context) {
	var equipe []EquipeProjetos
	id := c.Param("id")

	if equipe := h.DB.Raw("select eq.nome_equipe, pr.id_projeto, pr.nome_projeto from equipes as eq inner join projetos as pr on eq.id_equipe = pr.equipe_id where eq.id_equipe = ?", id).Scan(&equipe); equipe.Error != nil {
		c.AbortWithError(http.StatusNotFound, equipe.Error)
		return
	}

	c.JSON(http.StatusOK, &equipe)
}