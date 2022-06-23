package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)
type Tasks struct {
	ID_Task         uint 	`gorm:"primary_key" json:"id_task"`
	Descricao_Task  string 	`gorm:"type: varchar(100) not null" json:"descricao_task"`
	PessoaID  		uint	`json:"id_pessoa"`
	ProjetoID 		uint 	`json:"id_projeto"`
}

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	var task []Tasks

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)
}