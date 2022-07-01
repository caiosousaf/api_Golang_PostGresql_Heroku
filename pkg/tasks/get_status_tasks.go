package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetStatusTasks(c *gin.Context) {

	status := c.Param("status")

	var task []Task

	if result := h.DB.Raw("select * from tasks where status = ?", status).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)

}