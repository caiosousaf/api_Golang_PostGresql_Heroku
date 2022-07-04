package equipes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Membros struct {
	ID_Pessoa		int		`json:"id_pessoa"`
	Nome_Equipe 	string 	`json:"nome_equipe"`
	ID_Equipe		int		`json:"id_equipe"`
	Nome_Pessoa 	string 	`json:"nome_pessoa"`
	Funcao_Pessoa 	string 	`json:"funcao_pessoa"`

}

func (h handler) GetTeamMembers (c *gin.Context) {
	id := c.Param("id")

	var membros []Membros
	sql := `select eq.nome_equipe, id_equipe,pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa from equipes as eq inner join
	pessoas as pe on eq.id_equipe = pe.equipe_id where eq.id_equipe = ?`

	if membros := h.DB.Raw(sql, id).Scan(&membros); membros.Error != nil {
		c.AbortWithError(http.StatusNotFound, membros.Error)
		return
	}

	c.JSON(http.StatusOK, &membros)
}