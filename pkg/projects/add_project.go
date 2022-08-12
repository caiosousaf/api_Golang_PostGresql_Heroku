package projetos

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AddProjetoRequestBody struct {
	Nome_Projeto      string `gorm:"type: varchar(30) not null" json:"nome_projeto" example:"Projeto Teste"`
	Equipe_ID         int    `json:"equipe_id" example:"1"`
	Descricao_Projeto string `json:"descricao_projeto" example:"Projeto teste para exemplo do Swagger"`
	Prazo             int    `json:"prazo_entrega" example:"10"`
}

// @Summary POST a new Project
// @Description POST a new project. For the request to be met, the "nome_projeto", "equipe_id", "descricao_projeto" are required. The status already goes with a predefined value "A Fazer". the "prazo_entrega" is the number of days that the delivery time will be
// @Param		NewProject		body	string		true	"NewProject"
// @Accept json
// @Produce json
// @Success 201 {object} AddProjetoRequestBody
// @Failure 400 {array} models.Error400Create
// @Failure 404 {array} models.Error404Create
// @Tags Projects
// @Router /projetos [post]
func (h handler) AddProject(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error() ,
		})
		return
	}

	var t = body.Prazo
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)

	var projeto models.Projeto

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.Equipe_ID
	projeto.Status = "A Fazer" 
	projeto.Descricao_Projeto = body.Descricao_Projeto
	var count int


	// Verificando se já existe um projeto com o nome digitado
	if result := h.DB.Raw("select count(*) from projetos where nome_projeto = ?", body.Nome_Projeto).Scan(&count); result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
		})
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
		c.JSON(404, gin.H{
			"message": "Sla Cara. Descobre aí",
		})
		return
	}

	if result := h.DB.Model(&projeto).Where("id_projeto = ?", projeto.ID_Projeto).Update("prazo_entrega", data_limite.Format("2006-01-02")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

}
