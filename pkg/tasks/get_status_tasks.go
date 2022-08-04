package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GET status of tasks
// @Description GET All tasks with a specific status. "Em Andamento" or "Concluido"
// @Param		status		path	string		true		"Status"
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks/status/{status} [get]
func (h handler) GetStatusTasks(c *gin.Context) {

	status := c.Param("status")

	var task []Task

	if result := h.DB.Raw("select * from tasks where status = ?", status).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)

}