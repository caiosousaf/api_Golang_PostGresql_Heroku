package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetStatusList struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

func (h handler) GetStatusList(c *gin.Context) {
	var statuslist []GetStatusList

	if result := h.DB.Raw("select status, count(*) from projetos group by status").Scan(&statuslist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &statuslist)
}
