package equipes

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type EquipeProjetos struct {
	Nome_Equipe  		string 		`json:"nome_equipe"`
	ID_Projeto   		uint   		`json:"id_projeto"`
	Nome_Projeto 		string 		`json:"nome_projeto"`
	Status            	string     	`json:"status"`
    Data_Criacao      	string     	`json:"data_criacao"`
    Data_Conclusao    	string     	`json:"data_conclusao"`
    Prazo_Entrega     *time.Time 	`json:"prazo_entrega"`
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

	if equipe := h.DB.Raw(`select eq.nome_equipe, pr.id_projeto, pr.nome_projeto, pr.status, pr.descricao_projeto, pr.data_criacao, pr.data_criacao, pr.data_conclusao, pr.prazo_entrega 
	from equipes as eq 
	inner join projetos as pr on eq.id_equipe = pr.equipe_id where eq.id_equipe = ?`, id).Scan(&equipe); equipe.Error != nil {
		c.AbortWithError(http.StatusNotFound, equipe.Error)
		return
	}

	c.JSON(http.StatusOK, &equipe)
}