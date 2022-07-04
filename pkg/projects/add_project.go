package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	Nome_Projeto		string 				`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	Equipe_ID 			int					`json:"equipeid"`
	Status				string				`json:"status"`
	Descricao_Projeto	string				`json:"descricao_projeto"`
}

func (h handler) AddProject(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.Equipe_ID
	projeto.Status = "Não Iniciado"
	projeto.Descricao_Projeto = body.Descricao_Projeto

	if result := h.DB.Create(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &projeto)
}