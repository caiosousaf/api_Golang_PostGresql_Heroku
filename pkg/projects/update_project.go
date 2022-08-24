package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateProjetoRequestBody struct {
	Nome_Projeto      string `json:"nome_projeto"`
	Equipe_ID         int    `json:"equipe_id"`
	Descricao_Projeto string `json:"descricao_projeto"`
}

type UpdateStatusProject struct {
	Status string `json:"status"`
}

// PUT Project
// @Security bearerAuth
// @Summary PUT Project with ID
// @Description PUT a specific project. For the request to be met, the "nome_projeto" and "equipe_id" and "descricao_projeto" are required
// @Param        id   				path      	int  	true  	"Projeto ID"
// @Param		Project				body		string 	true 	"Project"
// @Accept json
// @Produce json
// @Success 200 {object} UpdateProjetoRequestBody
// @Failure 400,404 {string} string "error"
// @Tags Projects
// @Router /projetos/{id} [put]
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
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.DB.Save(&projeto)

	c.JSON(http.StatusOK, &projeto)
}

// @Summary PUT Status of a Project
// @Description PUT Status of a specific project. For the request to be met, the "status" are required
// @Param        id   						path      	int  	true  	"id"
// @Param		Status-Project				body		string 	true 	"Status-Project"
// @Accept json
// @Produce json
// @Success 200 {object} UpdateStatusProject
// @Failure 304	{string}	string "Not Modified"
// @Failure 400,404 {string} string "error"
// @Tags Projects
// @Router /projetos/{id}/status [put]
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

	if result := h.DB.Raw("update projetos set data_conclusao = current_date where status = 'Concluido' and id_projeto = ?", id).Scan(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &body)
}
