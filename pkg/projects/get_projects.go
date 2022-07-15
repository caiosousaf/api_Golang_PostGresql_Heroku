package projetos

import (
	"net/http"

	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)
type Projects struct {
	ID_Projeto 			uint 		`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 		string 		`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	EquipeID 			int 		`json:"equipe_id"`
	Nome_Equipe			string		`json:"nome_equipe"`
	Status				string		`json:"status"`
	Descricao_Projeto	string		`json:"descricao_projeto"`
	Data_Criacao		string		`json:"data_criacao"`
	Data_Conclusao		string		`json:"data_conclusao"`
}

type CountTeam struct {
	Count	int		`json:"count_team"`
}

func (h handler) GetProjects(c *gin.Context) {
	var projetos []Projects
	var status []GetStatusList
	var len []CountTeam

	if result := h.DB.Raw(`select pr.id_projeto, pr.nome_projeto, pr.equipe_id, eq.nome_equipe, pr.status, pr.descricao_projeto, 
	pr.data_criacao, pr.data_conclusao
	from projetos as pr inner join equipes as eq on pr.equipe_id = eq.id_equipe`).Scan(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw(`select status, count(*) from projetos group by status`).Scan(&status); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw(`select count(*) from equipes`).Scan(&len); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
   
	c.JSON(http.StatusOK, &projetos)
	c.JSON(http.StatusOK, &status)
	c.JSON(http.StatusOK, &len)
}
