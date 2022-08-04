package tasks

import (
	"net/http"
	"time"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Descricao_Task string `json:"descricao_task"`
	PessoaID       int    `json:"pessoa_id"`
	ProjetoID      int    `json:"projeto_id"`
	Prazo          int    `json:"prazo_entrega"`
	Prioridade     int    `json:"prioridade"`
}

// @Summary POST a new Task
// @Description POST a new task. For the request to be met, the "descricao_task", "pessoa_id", "projeto_id", "prazo_entrega(in days)", "prioridade" are required. The status already goes with a predefined value "Em Andamento". the "prazo_entrega" is the number of days that the delivery time will be
// @Param		NewTask		body	string		true	"NewTask"
// @Accept json
// @Produce json
// @Success 200 {object} AddTaskRequestBody
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks [post]
func (h handler) AddTask(c *gin.Context) {
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
	// Prazo em dias que será acrescentado a data atual para que a data_prazo seja retornada
	var t = body.Prazo
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Status = "Em Andamento"
	task.Prioridade = body.Prioridade


	var StatusCount int
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)
	err := c.ShouldBindJSON(&task)
	// Verifica se a pessoa em questão está na equipe atribuida ao projeto, se não estiver ele ira criar a tarefa

	if validationStatus := h.DB.Raw(`select count(*) from projetos where id_projeto = ?
	  and status = 'Em Andamento'`, body.ProjetoID).Scan(&StatusCount); validationStatus.Error != nil {
		c.AbortWithError(http.StatusBadRequest, validationStatus.Error)
		return
	}
	// se o StatusCount for > 0 então quer dizer que o projeto ainda está com o status "Em Andamento" se não estiver então ele não cria a tarefa
	if StatusCount > 0 {
		if result := h.DB.Create(&task); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		c.JSON(http.StatusCreated, &task)
		}	else {
			c.JSON(400, gin.H{
				"error": "Cannot create Task. Project is not under development: " + err.Error(),
			})
		}
	

	if result := h.DB.Model(&task).Where("id_task = ?", task.ID_Task).Update("prazo_entrega", data_limite.Format("2006-01-02")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

}