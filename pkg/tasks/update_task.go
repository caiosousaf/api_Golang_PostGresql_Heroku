package tasks

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

type UpdateTaskRequestBody struct {
	Descricao_Task  string 			`gorm:"type: varchar(100) not null" json:"descricao_task"`
	Pessoa			models.Pessoa 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pessoa"`
	Projeto			models.Projeto 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"projeto"`
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskRequestBody{}

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

	task.Descricao_Task = body.Descricao_Task
	task.Pessoa = body.Pessoa
	task.Projeto = body.Projeto

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}