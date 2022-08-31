package pessoas

import (
	"net/http"
	"strconv"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	ID_Task			uint	`json:"id_task"`
	Descricao_Task  string 	`json:"descricao_task"`
	Nivel			string	`json:"nivel"`
}

func (h handler) AddTaskPessoa(c *gin.Context) {
	id := c.Param("id")
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	var projeto int

	if result := h.DB.Raw(`select pr.id_projeto from projetos as pr
	inner join pessoas as pe
	on pe.equipe_id = pr.equipe_id
	where pe.id_pessoa = ? and pr.status = 'Em desenvolvimento'`, id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.ID_Task = body.ID_Task
	task.Descricao_Task = body.Descricao_Task
	task.Nivel = body.Nivel
	task.PessoaID = i
	task.ProjetoID = projeto
	task.Status = "A fazer"

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}