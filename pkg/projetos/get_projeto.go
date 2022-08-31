package projetos

import (
	"net/http"
	"time"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type BodyGetProjetos struct {
	ID_Projeto 			uint 			`json:"id_projeto"`
	Nome_Projeto 		string 			`json:"nome_projeto"`
	Descricao_Projeto 	string 			`json:"descricao_projeto"`
	Status 				string 			`json:"status"`
	DataInicio 			*time.Time 		`json:"data_inicio"`
	DataConclusao 		*time.Time		`json:"data_conclusao"`
	Equipe				models.Equipe	`json:"equipe"`
	Tasks				*[]models.Task 	`json:"tasks"`
}

func (h handler) GetProjeto(c *gin.Context) {
	id := c.Param("id")

	var projeto models.Projeto
	var pr BodyGetProjetos

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var equipe models.Equipe
	if result := h.DB.Where("id_equipe = ?", projeto.EquipeID).Find(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var tasks []models.Task
	if result := h.DB.Where("projeto_id = ?", projeto.ID_Projeto).Find(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pr.ID_Projeto = projeto.ID_Projeto
	pr.Nome_Projeto = projeto.Nome_Projeto
	pr.Status = projeto.Status
	pr.DataInicio = projeto.DataInicio
	pr.DataConclusao = projeto.DataConclusao
	pr.Equipe = equipe
	pr.Tasks = &tasks

	c.JSON(http.StatusOK, &pr)
}