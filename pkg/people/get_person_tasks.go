package pessoas

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type result struct {
	Nome_Pessoa    string     `json:"nome_pessoa"`
	Funcao_Pessoa  string     `json:"funcao_pessoa"`
	ID_Equipe	   int		  `json:"id_equipe"`
	Nome_Equipe    string     `json:"nome_equipe"`
	Nome_Projeto   string     `json:"nome_projeto"`
	ID_Task        int        `json:"id_task"`
	Descricao_Task string     `json:"descricao_task"`
	Projeto_ID     int        `json:"projeto_id"`
	Status         string     `json:"status"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     int        `json:"prioridade"`
}

 
// @Summary GET All Tasks of a specific Person
// @Description GET the tasks registered and assigned to a specific person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} result
// @Failure 400 {array} models.Error400Get
// @Failure 404 {array} models.Error404Get
// @Tags People
// @Router /pessoas/{id}/tasks [get]
func (h handler) GetTaskPerson(c *gin.Context) {
	var IdExist int
	id := c.Param("id")

	if result := h.DB.Raw("select count(*) from pessoas where id_pessoa = ?", id).Scan(&IdExist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	sql := `SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, eq.id_equipe, eq.nome_equipe, pr.nome_projeto,tk.id_task, tk.descricao_task, tk.projeto_id,
	 tk.status, tk.data_criacao, tk.data_conclusao, tk.prazo_entrega, tk.prioridade FROM
	 pessoas AS pe INNER JOIN equipes AS eq ON pe.equipe_id = eq.id_equipe INNER JOIN projetos AS pr ON pr.equipe_id = eq.id_equipe 
	 INNER JOIN tasks as tk ON tk.projeto_id = pr.id_projeto AND tk.pessoa_id = pe.id_pessoa WHERE pe.id_pessoa = ?`
	var result []result

	if IdExist == 1 {
		if result := h.DB.Raw(sql, id).Scan(&result); result.Error != nil {
			return
		}
	
		c.JSON(http.StatusOK, &result)
	} else {
		c.JSON(400, gin.H{
			"message": "Data not found with the passed parameters " ,
		})
		return
	}

}
