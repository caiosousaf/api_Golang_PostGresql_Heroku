package tasks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	ID_Task			uint		`json:"id_task"`
	Descricao_Task  string 		`json:"descricao_task"`
	PessoaID		int 		`json:"pessoa_id"`
	ProjetoID		int 		`json:"projeto_id"`
	Status			string		`json:"status"`
	Prazo			int			`json:"prazo_entrega"`
}

func (h handler) AddTask(c *gin.Context) {
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
	var t = body.Prazo
	var data_atual = time.Now() 
	data_limite := data_atual.AddDate(0,0,t)

	task.ID_Task = body.ID_Task
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Status = "Em Andamento"

	

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Model(&task).Where("id_task = ?", task.ID_Task).Update("prazo_entrega", data_limite.Format("2006-01-02")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	fmt.Println(t)

	c.JSON(http.StatusCreated, &task)
}