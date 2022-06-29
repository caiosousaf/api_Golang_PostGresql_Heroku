package tasks

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

type UpdateTaskRequestBody struct {
	ID_Task			uint			`json:"id_task"`
	Descricao_Task  string 			`json:"descricao_task"`
	PessoaID  		int				`json:"pessoa_id"`
	ProjetoID 		int 			`json:"projeto_id"`
	Status			int				`json:"status"`
}

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
	task.Status = body.Status

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}