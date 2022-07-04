package projetos

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
)

type UpdateProjetoRequestBody struct {

	Nome_Projeto		string 				`json:"nome_projeto"`
	Equipe_ID 			int					`json:"equipeid"`
	Status				string				`json:"status"`
	Descricao_Projeto	string				`json:"descricao_projeto"`
}

type UpdateStatusProject struct {
	Status string `json:"status"`
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
	projeto.Descricao_Projeto = body.Descricao_Projeto


	if result := h.DB.Raw(`update projetos set nome_projeto = ?, equipe_id = ?, descricao_projeto = ? where id_projeto = ?`, projeto.Nome_Projeto, projeto.EquipeID, projeto.Descricao_Projeto, id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
	}

	h.DB.Save(&projeto)

	c.JSON(http.StatusOK, &projeto)
}

func (h handler) UpdateStatusProject(c *gin.Context) {
	id := c.Param("id")
	body := UpdateStatusProject{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Raw("update projetos set status = ? where id_projeto = ?", body.Status, id).Scan(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &body)
}