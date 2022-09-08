package tasks

import (
	"database/sql"
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

func ListarTasks(c *gin.Context) {
	fmt.Println("Tentando listar todos as tasks")
	if tasks, err := tasks.ListarTasks(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(200, tasks)
	}
}