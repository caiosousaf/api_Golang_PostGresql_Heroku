package equipes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (h handler) GetTeamsMembers (c *gin.Context) {

	var membros []Membros
	sql := `select eq.nome_equipe, id_equipe,pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa from equipes as eq inner join
	pessoas as pe on eq.id_equipe = pe.equipe_id`

	if membros := h.DB.Raw(sql).Scan(&membros); membros.Error != nil {
		c.AbortWithError(http.StatusNotFound, membros.Error)
		return
	}

	c.JSON(http.StatusOK, &membros)
}