package tasks

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Descricao_Task  string 			`gorm:"type: varchar(100) not null" json:"descricao"`
	Pessoa			models.Pessoa 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pessoa"`
	Projeto			models.Projeto 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"projeto"`
}

func (h handler) AddTask(c *gin.Context) {
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Descricao_Task = body.Descricao_Task
	task.Pessoa = body.Pessoa
	task.Projeto = body.Projeto

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}