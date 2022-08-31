package tasks

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Tasks struct {
	ID_Task 		uint 	`json:"id_task"`
	Descricao_Task 	string 	`json:"descricao_task"`
	Nivel			string	`json:"nivel"`
	Status 			string 	`json:"status"`
	PessoaID 		int 	`json:"pessoa_id"`
	Nome_Pessoa 	string 	`json:"nome_pessoa"`
	ProjetoID 		int 	`json:"projeto_id"`
	Nome_Projeto 	string 	`json:"nome_projeto"`
}

func (h handler) GetTasks(c *gin.Context) {
	var tasks []Tasks
	
	sql := `select tk.id_task, tk.descricao_task, tk.nivel, tk.status, tk.pessoa_id, pe.nome_pessoa, 
	tk.projeto_id, pr.nome_projeto from tasks as tk inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa 
	inner join projetos as pr on tk.projeto_id = pr.id_projeto order by id_task`

	if tasks := h.DB.Raw(sql).Scan(&tasks); tasks.Error != nil {
		c.AbortWithError(http.StatusNotFound, tasks.Error)
		return
	}
	
	c.JSON(http.StatusOK, &tasks)
}
