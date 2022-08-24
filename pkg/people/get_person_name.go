package pessoas

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary GET a specific Person by Name
// @Description GET a specific person by Name
// @Param		person				query			string				false		"name_person"
// @Param		person_function		query			string				false		"funcao_person"	Enums(Back-End, Front-End, Analista, Tester)
// @Accept json
// @Produce json
// @Success 200 {array} GetPessoa
// @Failure 400,404 {string} string "error"
// @Tags People
// @Router /pessoas/filtros/ [get]
func (h handler) GetPersonName(c *gin.Context) {
	name := c.Query("person")
	function_pe := c.Query("person_function")

	NameOrFunc := "nome_pessoa"

	if function_pe == "Back-End" || function_pe == "Front-End" || function_pe == "Analista" || function_pe == "Tester"{
		NameOrFunc = "funcao_pessoa"
	}
	if name == "" {
		name = function_pe
	}

	var pessoa []GetPessoa
	fmt.Println(name)
	if pessoa := h.DB.Raw(`select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao 
	from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe where `+ NameOrFunc + ` ilike ?` , name).Scan(&pessoa); pessoa.Error != nil {
		c.AbortWithError(http.StatusNotFound, pessoa.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}