package projetos

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

type UpdateProjetoRequestBody struct {
	ID_Projeto 			uint 				`json:"id_projeto"`
	Nome_Projeto		string 				`json:"nome_projeto"`
	Equipe_ID 			int					`json:"equipeid"`
	Status				string				`json:"status"`
	Descricao_Projeto	string				`json:"descricao_projeto"`
}

func (h handler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	body := UpdateProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.Equipe_ID
	projeto.Status = body.Status
	projeto.Descricao_Projeto = body.Descricao_Projeto
	projeto.ID_Projeto = body.ID_Projeto

	if result := h.DB.Raw(`update projetos set nome_projeto = ?, equipe_id = ?, status = ?, descricao_projeto = ? where id_projeto = ?`, projeto.Nome_Projeto, projeto.EquipeID, projeto.Status, projeto.Descricao_Projeto, projeto.ID_Projeto).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
	}

	h.DB.Save(&projeto)

	c.JSON(http.StatusOK, &projeto)
}