package projetos

import (
	"net/http"
	"time"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	Nome_Projeto		string 				`gorm:"type: varchar(30) not null" json:"nome_projeto"`
	Equipe_ID 			int					`json:"equipe_id"`
	Status				string				`json:"status"`
	Descricao_Projeto	string				`json:"descricao_projeto"`
	Data_Criacao		string				`json:"data_criacao"`
}

func (h handler) AddProject(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto
	dt := time.Now()

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.Equipe_ID
	projeto.Status = "Em Andamento"
	projeto.Descricao_Projeto = body.Descricao_Projeto
	projeto.Data_Criacao = dt.Format("02-01-2006")


	if result := h.DB.Create(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &projeto)
}