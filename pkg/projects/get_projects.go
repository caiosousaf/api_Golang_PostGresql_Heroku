package projetos

import (
	"net/http"
	"time"

	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Projects struct {
	ID_Projeto        uint       `gorm:"primary_key" json:"id_projeto" example:"58"`
	Nome_Projeto      string     `gorm:"type: varchar(30) not null" json:"nome_projeto" example:"Nome"`
	Descricao_Projeto string     `json:"descricao_projeto" example:"Descricao"`
	EquipeID          int        `json:"equipe_id" example:"2"`
	Nome_Equipe       string     `json:"nome_equipe" example:"Cariri Inovação"`
	Status            string     `json:"status" example:"Concluido"`
	Data_Criacao      string     `json:"data_criacao" example:"2022-07-25"`
	Data_Conclusao    string     `json:"data_conclusao" example:""`
	Prazo_Entrega     *time.Time `json:"prazo_entrega" example:"2022-07-25"`
}

// @Security bearerAuth
// Get Projects
// @Summary Get All Projects
// @Description Get list all project
// @Accept json
// @Produce json
// @Success 200 {array} Projects
// @Failure 404 {string} string "error"
// @Tags Projects
// @Router /projetos [get]
func (h handler) GetProjects(c *gin.Context) {
	var projetos []Projects

	if result := h.DB.Raw(`select pr.id_projeto, pr.nome_projeto, pr.equipe_id, eq.nome_equipe, pr.status, pr.descricao_projeto, 
	pr.data_criacao, pr.data_conclusao, pr.prazo_entrega
	from projetos as pr inner join equipes as eq on pr.equipe_id = eq.id_equipe`).Scan(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projetos)
}
