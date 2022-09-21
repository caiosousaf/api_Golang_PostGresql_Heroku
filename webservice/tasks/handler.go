package tasks

import (
	"gerenciadorDeProjetos/domain/tasks"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary POST a new Task
// @Description POST a new task. For the request to be met, the "descricao_task", "pessoa_id", "projeto_id", "prazo_entrega(in days)", "prioridade" are required. The status already goes with a predefined value "A Fazer". the "prazo_entrega" is the number of days that the delivery time will be
// @Param		NewTask		body	string		true	"NewTask"
// @Accept json
// @Produce json
// @Success 201 {object} modelApresentacao.ReqTaskApresent "OK"
// @Failure 401,400 {array} utils.ResError
// @Tags Tasks
// @Router /tasks [post]
func NovaTask(c *gin.Context) {

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

// @Security bearerAuth
// Get Tasks
// @Summary Get All Tasks
// @Description Get list all task
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTasks "OK"
// @Failure 401,404 {array} utils.ResError
// @Tags Tasks
// @Router /tasks [get]
func ListarTasks(c *gin.Context) {

	if tasks, err := tasks.ListarTasks(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// @Security bearerAuth
// @Summary Get a specific Task
// @Description Get a specific task with id
// @Param	id		path	int		true	"Task ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTasks "OK"
// @Failure 401,404 {array} utils.ResError
// @Tags Tasks
// @Router /tasks/{id} [get]
func ListarTask(c *gin.Context) {

	id := c.Param("id")
	if tasks, err := tasks.ListarTask(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Task does not exist", 404))
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// @Security bearerAuth
// @Summary GET status of tasks
// @Description GET All tasks with a specific status. "Em Andamento" or "Concluido"
// @Param		status		path	string		true		"Status"	Enums(A Fazer,Em Andamento,Em Teste,Concluido)
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTasks "OK"
// @Failure 401,404 {array} utils.ResError
// @Tags Tasks
// @Router /tasks/status/{status} [get]
func ListarStatusTasks(c *gin.Context) {

	status := c.Param("status")
	if tasks, err := tasks.ListarStatusTasks(status); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Status does not exist", 404))
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// @Security bearerAuth
// @Summary PUT Task
// @Description PUT a specific task. For the request to be met, the "descricao_task" and "pessoa_id" and "projeto_id" and "prioridade" are required.
// @Param        id   				path      	int  	true  	"Task ID"
// @Param		Task				body		string 	true 	"PUT Task"
// @Accept json
// @Produce json
// @Success 200 {object} tasks.ReqUpdateTaskData "OK"
// @Failure 401,400 {array} utils.ResError
// @Tags Tasks
// @Router /tasks/{id} [put]
func AtualizarTask(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
	}

	if res, err := tasks.AtualizarTask(id, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @Security bearerAuth
// @Summary PUT Status of a Task
// @Description PUT Status of a specific Task. For the request to be met, the "status" are required
// @Param        id   				path      	int  	true  	"Task ID"
// @Param		Status				body		string 	true 	"Status"
// @Accept json
// @Produce json
// @Success 200 {object} tasks.ReqUpdateStatusTask
// @Failure 401,400,404 {array} utils.ResError
// @Tags Tasks
// @Router /tasks/{id}/status [put]
func AtualizarStatusTask(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly", "err": err.Error(),
		})
		return
	}

	if res, err := tasks.AtualizarStatusTask(id, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @Security bearerAuth
// @Summary DELETE a Task
// @Description DELETE a Task with id
// @Param		id		path	int		true		"Task_ID"
// @Accept json
// @Produce json
// @Success 200 {array} utils.ResOk
// @Failure 401,404 {array} utils.ResError
// @Tags Tasks
// @Router /tasks/{id} [delete]
func DeletarTask(c *gin.Context) {
	id := c.Param("id")

	err := tasks.DeletarTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Task deleted successfully", 200))
	}
}
