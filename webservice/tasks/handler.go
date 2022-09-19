package tasks

import (
	"fmt"
	"gerenciadorDeProjetos/domain/tasks"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(), "Team does not exist", 400))
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func ListarTasks(c *gin.Context) {
	fmt.Println("Tentando listar todos as tasks")
	if tasks, err := tasks.ListarTasks(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func ListarTask(c *gin.Context) {
	fmt.Println("Tentando listar uma task")
	id := c.Param("id")
	if tasks, err := tasks.ListarTask(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Task does not exist", 404))
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func ListarStatusTasks(c *gin.Context) {
	fmt.Println("Tentando listar todas as tarefas com um status especifico")
	status := c.Param("status")
	if tasks, err := tasks.ListarStatusTasks(status); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Status does not exist", 404))
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func AtualizarTask(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar uma task")

	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
	}

	if res, err := tasks.AtualizarTask(id, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func AtualizarStatusTask(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar status de uma task")

	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly", "err": err.Error(),
		})
		return
	}

	if res, err := tasks.AtualizarStatusTask(id, &req); err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func DeletarTask(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma task")

	err := tasks.DeletarTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Task deleted successfully", 200))
	}
}
