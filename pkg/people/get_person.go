package pessoas

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPessoa struct {
	ID_Pessoa        uint   `json:"id_pessoa"`
	Nome_Pessoa      string `json:"nome_pessoa"`
	Funcao_Pessoa    string `json:"funcao_pessoa"`
	EquipeID         int    `json:"equipe_id"`
	Nome_Equipe      string `json:"nome_equipe"`
	Data_Contratacao string `json:"data_contratacao"`
}

// @Summary GET a specific Person
// @Description GET a specific person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} GetPessoa
// @Failure 400 {array} models.Error400Get
// @Failure 404 {array} models.Error404Get
// @Tags People
// @Router /pessoas/{id} [get]
func (h handler) GetPerson(c *gin.Context) {
	id := c.Param("id")
	var IdExist int

	var pessoa GetPessoa

	if result := h.DB.Raw("select count(*) from pessoas where id_pessoa = ?", id).Scan(&IdExist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if IdExist == 1 {
		if pessoa := h.DB.Raw(`select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao
		from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe where id_pessoa = ?`, id).Scan(&pessoa); pessoa.Error != nil {
		c.JSON(404, gin.H{
			"message": "Loss of contact with the database" ,
		})
		return
		}

		c.JSON(http.StatusOK, &pessoa)
	} else {
		c.JSON(400, gin.H{
			"message": "Data not found with the passed parameters " ,
		})
		return
	}
	
}
