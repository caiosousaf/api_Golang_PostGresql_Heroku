package tasks

import (
	"net/http"


	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateStatusTaskRequestBody struct {
	Status			string				`json:"status"`
}

func (h handler) UpdateStatusTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateStatusTaskRequestBody{}

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

	task.Status	= body.Status


	if tasks := h.DB.Raw("update tasks set status = ? where id_task = ?", task.Status, id).Scan(&task); tasks.Error != nil {
		c.AbortWithError(http.StatusNotFound, tasks.Error)
		return
	}

	if tasks := h.DB.Raw("update tasks set data_conclusao = current_date where status = 'Concluido' and id_task = ?", id).Scan(&task); tasks.Error != nil {
		c.AbortWithError(http.StatusNotFound, tasks.Error)
		return
	}

	c.JSON(http.StatusOK, &task)
}