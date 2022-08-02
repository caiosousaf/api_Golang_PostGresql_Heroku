package tasks

import (
	"net/http"
	"time"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	ID_Task        uint   `json:"id_task"`
	Descricao_Task string `json:"descricao_task"`
	PessoaID       int    `json:"pessoa_id"`
	ProjetoID      int    `json:"projeto_id"`
	Status         string `json:"status"`
	Prazo          int    `json:"prazo_entrega"`
	Prioridade     int    `json:"prioridade"`
}

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
	task.ID_Task = body.ID_Task
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Status = "Em Andamento"
	task.Prioridade = body.Prioridade

	var count int
	var StatusCount int
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)
	err := c.ShouldBindJSON(&task)
	// Verifica se a pessoa em questão está na equipe atribuida ao projeto, se não estiver ele ira criar a tarefa
	if result := h.DB.Raw(`select count(*) from pessoas as pe inner join projetos as pr on pe.equipe_id = pr.equipe_id where pe.id_pessoa = ? and
	 pr.id_projeto = ?`, body.PessoaID, body.ProjetoID).Scan(&count); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if validationStatus := h.DB.Raw(`select count(*) from projetos where id_projeto = ?
	  and status = 'Em Andamento'`, body.ProjetoID).Scan(&StatusCount); validationStatus.Error != nil {
		c.AbortWithError(http.StatusBadRequest, validationStatus.Error)
		return
	}
	// Se o count for > ) então quer dizer que a pessoa está na equipe em que o projeto foi atribuido
	
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