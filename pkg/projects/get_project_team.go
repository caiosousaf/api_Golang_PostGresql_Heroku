package projetos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Equipes struct {
	ID_projeto		  uint 		`json:"id_projeto"`
	ID_Equipe         uint   	`json:"id_equipe"`
    Nome_Equipe       string 	`json:"nome_equipe"`
}

func (h handler) GetProjectTeam(c * gin.Context){
	id := c.Param("id")

	var equipes []Equipes
	
	sql := "select eq.nome.equipe from equipes as eq inner join equipes as pe on eq.id_projeto = pe.projeto_id where eq.id_equipe = ?"

	if equipes := h.DB.Raw(sql, id).Scan(&equipes); equipes.Error != nil{
		c.AbortWithError(http.StatusNotFound, equipes.Error)
		return
	}
	c.JSON(http.StatusOK, &equipes)
}