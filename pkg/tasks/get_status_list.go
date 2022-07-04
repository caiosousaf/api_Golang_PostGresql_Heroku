package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetStatusList struct {
	Status		string	`json:"status"`
	Quantidade	int		`json:"quantidade"`
}

func (h handler) GetStatusList(c *gin.Context) {
	var statuslist []GetStatusList

	if result := h.DB.Raw("select * from projetos order by id_projeto").Scan(&statuslist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	c.JSON(http.StatusOK, &statuslist)
}