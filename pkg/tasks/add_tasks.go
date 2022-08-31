package tasks

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Descricao_Task  string 	`json:"descricao_task"`
	Nivel			string	`json:"nivel"`
	PessoaID		int 	`json:"pessoa_id"`
	ProjetoID		int 	`json:"projeto_id"`
}

func (h handler) AddTask(c *gin.Context) {
	
	body := AddTaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Descricao_Task = body.Descricao_Task
	task.Nivel = body.Nivel
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Status = "A fazer"

	// Pegando ID da equipe em que o Projeto está cadastrado
	var equipe int
	if result := h.DB.Raw("select equipe_id from projetos where id_projeto = ? and equipe_id is not null", task.ProjetoID).Scan(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	/* Verifica se o projeto está cadastrado em alguma equipe */
	if equipe == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error":"O projeto informado não está atribuído a nenhuma equipe"})
		return
	}

	// A variável checkE vai armazenar o resultado do count que o sql verifica_equipe vai retornar
	var checkE int

	/* verifica_equipe vai retornar um count que vai indicar se a pessoa que pegou essa task está realmente na equipe
	responsável pelo projeto, e se o projeto em questão está com o status 'Em desenvolvimento', pois uma task só
	pode ser cadastrada em projetos em desenvolvimento */
	verifica_equipe := `select count(pe.id_pessoa) from pessoas as pe
	inner join projetos as pr on pe.equipe_id = pr.equipe_id
	where pe.id_pessoa = ? and pr.equipe_id = ?`

	if result := h.DB.Raw(verifica_equipe, body.PessoaID, equipe).Scan(&checkE); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if(checkE > 0){

		var checkS int
		verifica_status := "select count(id_projeto) from projetos where id_projeto = ? and status = 'Em desenvolvimento'"

		if result := h.DB.Raw(verifica_status, body.ProjetoID).Scan(&checkS); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		if(checkS > 0){
			if result := h.DB.Create(&task); result.Error != nil {
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
		
			c.JSON(http.StatusCreated, &task)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":"Tasks só podem ser cadastradas em projetos que estão 'Em desenvolvimento'."})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error":"Esta pessoa não está na mesma equipe que o projeto."})
		return
	}
	
}