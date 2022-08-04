package tasks

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateTaskRequestBody struct {
	Descricao_Task string `json:"descricao_task"`
	PessoaID       int    `json:"pessoa_id"`
	ProjetoID      int    `json:"projeto_id"`
	Prioridade     int    `json:"prioridade"`
}

// @Summary PUT Task 
// @Description PUT a specific task. For the request to be met, the "descricao_task" and "pessoa_id" and "projeto_id" and "prioridade" are required.
// @Param        id   				path      	int  	true  	"Task ID"
// @Param		Task				body		string 	true 	"PUT Task"
// @Accept json
// @Produce json
// @Success 200 {object} UpdateTaskRequestBody
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks/{id} [put]
func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Prioridade = body.Prioridade

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}