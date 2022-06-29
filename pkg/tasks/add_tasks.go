package tasks

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	ID_Task			uint	`json:"id_task"`
	Descricao_Task  string 	`json:"descricao_task"`
	PessoaID		int 	`json:"pessoa_id"`
	ProjetoID		int 	`json:"projeto_id"`
}

func (h handler) AddTask(c *gin.Context) {
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.ID_Task = body.ID_Task
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}