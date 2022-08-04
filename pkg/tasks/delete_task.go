package tasks

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Summary DELETE a Task
// @Description DELETE a Task with id
// @Param		id		path	int		true		"Task_ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks/{id} [delete]
func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&task)

	c.Status(http.StatusOK)
}