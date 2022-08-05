package tasks

import (
	"net/http"
	"time"

	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID_Task        uint       `json:"id_task"`
	Descricao_Task string     `json:"descricao_task"`
	PessoaID       int        `json:"pessoa_id"`
	Nome_Pessoa    string     `json:"nome_pessoa"`
	ProjetoID      int        `json:"projeto_id"`
	Nome_Projeto   string     `json:"nome_projeto"`
	Status         string     `json:"status" enums:"Em Andamento, Concluido"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     int        `json:"prioridade"`
}

// Get Tasks
// @Summary Get All Tasks
// @Description Get list all task
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks [get]
func (h handler) GetTasks(c *gin.Context) {
	var tasks []Task
	sql := "select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade from tasks as tk inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa inner join projetos as pr on tk.projeto_id = pr.id_projeto order by id_task"
	if tasks := h.DB.Raw(sql).Scan(&tasks); tasks.Error != nil {
		c.AbortWithError(http.StatusNotFound, tasks.Error)
		return
	}

	c.JSON(http.StatusOK, &tasks)
}