package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksProjeto struct {
	ID_Projeto     uint   `json:"id_projeto"`
	Nome_Projeto   string `json:"nome_projeto"`
	Nome_Equipe    string `json:"nome_equipe"`
	ID_Task        string `json:"id_task"`
	Descricao_Task string `json:"descricao_task"`
	Status         string `json:"status"`
	Nome_Pessoa    string `json:"nome_pessoa"`
}

func (h handler) GetProjectTasks(c *gin.Context) {
	var tasks []TasksProjeto

	id := c.Param("id")
	sql := `select pr.id_projeto, pr.nome_projeto, eq.nome_equipe,tk.id_task, tk.descricao_task, tk.status, 
	pe.nome_pessoa from 
	projetos as pr inner join tasks as tk on pr.id_projeto = tk.projeto_id inner join
	equipes as eq on pr.equipe_id = eq.id_equipe inner join
	pessoas as pe on pe.id_pessoa = tk.pessoa_id where id_projeto = ?`

	if tasks := h.DB.Raw(sql, id).Scan(&tasks); tasks.Error != nil {
		c.AbortWithError(http.StatusNotFound, tasks.Error)
		return
	}

	c.JSON(http.StatusOK, &tasks)
}