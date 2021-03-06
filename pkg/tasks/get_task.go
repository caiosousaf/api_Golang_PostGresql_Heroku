package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)


func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	var task []Task

	if result := h.DB.Raw(`select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id,
	pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao ,tk.prazo_entrega, tk.prioridade from tasks as tk inner join pessoas as pe on tk.pessoa_id
	= pe.id_pessoa inner join projetos as pr on tk.projeto_id = pr.id_projeto where tk.id_task = ?`, id).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)
}