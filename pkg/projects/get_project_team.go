package projetos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Teams struct {
	ID_project	uint 	`json:"id_project"`
	ID_Team		uint	`json:"id_equipe"`
	NameTeam 	string 	`json:"nameTeam"`
}

func (h handler) GetProjectTeam(c * gin.Context){
	id := c.Param("id")

	var teams []Teams
	
	sql := "select eq.nameteam from teams as eq inner join teams as pe on eq.id_project = pe.project_id where eq.id_team = ?"

	if teams := h.DB.Raw(sql, id).Scan(&teams); teams.Error != nil{
		c.AbortWithError(http.StatusNotFound, teams.Error)
		return
	}
	c.JSON(http.StatusOK, &teams)
}