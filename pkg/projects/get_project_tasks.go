package projetos

import (
	"net/http"
	"time"

	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type TasksProjeto struct {
	ID_Projeto     uint       `json:"id_projeto"`
	Nome_Projeto   string     `json:"nome_projeto"`
	Nome_Equipe    string     `json:"nome_equipe"`
	ID_Task        string     `json:"id_task"`
	Descricao_Task string     `json:"descricao_task"`
	Pessoa_ID      int        `json:"pessoa_id"`
	Nome_Pessoa    string     `json:"nome_pessoa"`
	Status         string     `json:"status"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     int        `json:"prioridade"`
}

func (h handler) GetProjectTasks(c *gin.Context) {
	var tasks []TasksProjeto

	id := c.Param("id")
	sql := `select tk.*, pr.id_projeto, pr.nome_projeto, eq.nome_equipe,
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
