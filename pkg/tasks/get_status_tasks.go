package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// @Summary GET status of tasks
// @Description GET All tasks with a specific status. "Em Andamento" or "Concluido"
// @Param		status		path	string		true		"Status"	Enums(Em Andamento, Concluido)	
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks/status/{status} [get]
func (h handler) GetStatusTasks(c *gin.Context) {

	status := c.Param("status")

	var task []Task

	if result := h.DB.Raw(`select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status,
	 tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade
	 from tasks as tk inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa inner join projetos as pr on tk.projeto_id = pr.id_projeto 
	 where tk.status = ? order by id_task`, status).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)

}