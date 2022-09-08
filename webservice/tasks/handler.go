package tasks

import (
	"fmt"
	"gerenciadorDeProjetos/domain/tasks"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NovaTask(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova task")
	req := modelApresentacao.ReqTaskApresent{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "error": err.Error(),
		})
		return
	}

	if res, err := tasks.NovaTask(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}