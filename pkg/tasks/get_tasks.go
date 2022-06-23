package tasks

import (
	"net/http"

	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
	"github.com/gin-gonic/gin"
)
type Task struct {
	ID_Task         uint 	`gorm:"primary_key" json:"id_task"`
	Descricao_Task  string 	`gorm:"type: varchar(100) not null" json:"descricao_task"`
	PessoaID  		uint	`json:"id_pessoa"`
	ProjetoID 		uint 	`json:"id_projeto"`
}

func (h handler) GetTasks(c *gin.Context) {
	var tasks []Task

	if result := h.DB.Find(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	c.JSON(http.StatusOK, &tasks)
}
