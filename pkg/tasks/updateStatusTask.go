package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdateStatusTask(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")


	var task []Task

	if result := h.DB.Raw("update tasks set status = ? where id_task = ?", status, id).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)

} 