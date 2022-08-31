package tasks

import (
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdateTaskRequestBody struct {
	Descricao_Task  string 	`json:"descricao_task"`
	Nivel			string	`json:"nivel"`
	PessoaID		string 	`json:"pessoa_id"`
	ProjetoID		string 	`json:"projeto_id"`
}

type UpdateStatus struct {
	Status string `json:"status"`
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

	prId, err := strconv.Atoi(body.ProjetoID)
	peId, err2 := strconv.Atoi(body.PessoaID)

	if err == nil && err2 == nil{
		task.Descricao_Task = body.Descricao_Task
		task.Nivel = body.Nivel
		task.PessoaID = peId
		task.ProjetoID = prId
	} else{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}

func (h handler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	body := UpdateStatus{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Raw("update tasks set status = ? where id_task = ?", body.Status, id).Scan(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &body)
}