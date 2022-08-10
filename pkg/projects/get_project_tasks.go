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
	ID_Task        int     	  `json:"id_task"`
	Descricao_Task string     `json:"descricao_task"`
	Pessoa_ID      int        `json:"pessoa_id"`
	Nome_Pessoa    string     `json:"nome_pessoa"`
	Status         string     `json:"status"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     int        `json:"prioridade"`
}

// Get Tasks of Project
// @Summary Get Tasks of Project with Param ID
// @Description GET all tasks of a project with ID_Projeto specific
// @Param        id   path      int  true  "Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} TasksProjeto
// @Failure 400 {array} models.Error400Get
// @Failure 404 {array} models.Error404Message
// @Tags Projects
// @Router /projetos/{id}/tasks [get]
func (h handler) GetProjectTasks(c *gin.Context) {
	var tasks []TasksProjeto
	var CheckProjectid int

	id := c.Param("id")
	sql := `select tk.*, pr.id_projeto, pr.nome_projeto, eq.nome_equipe,
	pe.nome_pessoa from 
	projetos as pr inner join tasks as tk on pr.id_projeto = tk.projeto_id inner join
	equipes as eq on pr.equipe_id = eq.id_equipe inner join
	pessoas as pe on pe.id_pessoa = tk.pessoa_id where id_projeto = ?`

	if result := h.DB.Raw("select count(id_projeto) from projetos where id_projeto = ?", id).Scan(&CheckProjectid); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	if CheckProjectid > 0 {
		if tasks := h.DB.Raw(sql, id).Scan(&tasks); tasks.Error != nil {
			c.AbortWithError(http.StatusNotFound, tasks.Error)
			return
		}
		c.JSON(http.StatusOK, &tasks)
		
	} else {
		c.JSON(400, gin.H{
			"message": "Data not found with the passed parameters" ,
		})
		return
	}

	
}
