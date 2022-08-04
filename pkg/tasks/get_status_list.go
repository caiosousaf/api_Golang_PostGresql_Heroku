package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetStatusList struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}


// @Summary Get a list of status of All Tasks
// @Description Get list of status of all tasks registered
// @Accept json
// @Produce json
// @Success 200 {array} GetStatusList
// @Failure 400,404 {string} string "error"
// @Tags Tasks
// @Router /tasks/list [get]
func (h handler) GetStatusList(c *gin.Context) {
	var statuslist []GetStatusList

	if result := h.DB.Raw("select status, count(*) from tasks group by status").Scan(&statuslist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &statuslist)
}