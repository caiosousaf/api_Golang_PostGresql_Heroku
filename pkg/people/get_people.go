package pessoas

import (
	"net/http"

	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPeople(c *gin.Context) {
	var pessoas []GetPessoa

	if pessoas := h.DB.Raw(`select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao
	from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe`).Scan(&pessoas); pessoas.Error != nil {
		c.AbortWithError(http.StatusNotFound, pessoas.Error)
		return
	}
	c.JSON(http.StatusOK, &pessoas)
}
