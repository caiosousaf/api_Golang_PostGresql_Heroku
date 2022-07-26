package projetos

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AddProjetoRequestBody struct {
	Nome_Projeto      string `gorm:"type: varchar(30) not null" json:"nome_projeto"`
	Equipe_ID         int    `json:"equipe_id"`
	Status            string `json:"status"`
	Descricao_Projeto string `json:"descricao_projeto"`
	Prazo             int    `json:"prazo_entrega"`
}

func (h handler) AddProject(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var t = body.Prazo
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)

	var projeto models.Projeto

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.Equipe_ID
	projeto.Status = "Em Andamento"
	projeto.Descricao_Projeto = body.Descricao_Projeto
	var count int

	err := c.ShouldBindJSON(&projeto)
	// Verificando se já existe um projeto com o nome digitado
	if result := h.DB.Raw("select count(*) from projetos where nome_projeto = ?", body.Nome_Projeto).Scan(&count); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Se não existir nenhum projeto com esse nome ele cria um novo projeto com sucesso
	if count == 0 {
		if result := h.DB.Create(&projeto); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		c.JSON(http.StatusCreated, &projeto)
	} else {
		c.JSON(400, gin.H{
			"error": "Cannot create Project. already existing project: " + err.Error(),
		})
	}

	if result := h.DB.Model(&projeto).Where("id_projeto = ?", projeto.ID_Projeto).Update("prazo_entrega", data_limite.Format("2006-01-02")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

}
