package equipes

import (
	"net/http"

	equipe "gerenciadorDeProjetos/domain/equipes"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary POST a new Team
// @Description POST a new Team. For the request to be met, the "nome_equipe" are required.
// @Param		Team		body	string		true	"NewTeam"
// @Accept json
// @Produce json
// @Success 201 {object} equipes.Equipe "OK"
// @Failure 401,400 {array} utils.ResError
// @Tags Teams
// @Router /equipes [post]
func novaEquipe(c *gin.Context) {

	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
			"Could not create. Parameters were not passed correctly", 400))
		return
	}

	if res, err := equipe.NovaEquipe(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Get Teams
// @Security bearerAuth
// @Summary Get All Teams
// @Description Returns all registered teams and all their members they are assigned to
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqEquipe "OK"
// @Failure 401,404 {string} string "error"
// @Tags Teams
// @Router /equipes [get]
func listarEquipes(c *gin.Context) {

	if equipes, err := equipe.ListarEquipes(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, equipes)
	}
}

// @Security bearerAuth
// @Summary Get Specific Team
// @Description Returns a team, all their members, all projects and all tasks they are assigned to
// @Param		id		path	int		true	"id_Team"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqEquipe "OK"
// @Failure 401,404 {array} utils.ResError
// @Tags Teams
// @Router /equipes/{id} [get]
func buscarEquipe(c *gin.Context) {
	id := c.Param("id")

	if equipe, err := equipe.BuscarEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

// @Security bearerAuth
// @Summary Get Members of a specific Team
// @Description GET all members of a specific Team with ID
// @Param        id   path      int  true  "Team ID"
// @Accept json
// @Produce json
// @Success 200 {array} pessoas.ReqMembros "OK"
// @Failure 401,404,204 {array} utils.ResError
// @Tags Teams
// @Router /equipes/{id}/membros [get]
func buscarMembrosDeEquipe(c *gin.Context) {
	id := c.Param("id")

	if equipe, err := equipe.BuscarMembrosDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

// @Security bearerAuth
// @Summary Get Projects of a specific Team
// @Description GET all projects of a specific Team with ID
// @Param        id   path      int  true  "Team ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqEquipeProjetos "OK"
// @Failure 401,404,204 {array} utils.ResError
// @Tags Teams
// @Router /equipes/{id}/projetos [get]
func buscarProjetosDeEquipe(c *gin.Context) {
	id := c.Param("id")

	if equipe, err := equipe.BuscarProjetosDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

// @Security bearerAuth
// @Summary Get Tasks of a specific Team
// @Description GET all tasks of a specific Team with ID
// @Param        id   path      int  true  "Team ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTasksbyTeam "OK"
// @Failure 401,404,204 {array} utils.ResError
// @Tags Teams
// @Router /equipes/{id}/tasks [get]
func buscarTasksDeEquipe(c *gin.Context) {
	id := c.Param("id")

	if equipe, err := equipe.BuscarTasksDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

// @Security bearerAuth
// @Summary DELETE a Team
// @Description DELETE a Team
// @Param		id		path	int		true		"Team_ID"
// @Accept json
// @Produce json
// @Success 200 {array} utils.ResOk
// @Failure 401,404 {array} utils.ResError
// @Tags Teams
// @Router /equipes/{id} [delete]
func deletarEquipe(c *gin.Context) {
	id := c.Param("id")

	if err := equipe.DeletarEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Team deleted successfully", 200))
	}
}

// PUT Team
// @Security bearerAuth
// @Summary PUT Team with ID
// @Description PUT a specific Team. For the request to be met, the "id_equipe" and "nome_equipe" are required
// @Param        id   				path      	int  	true  	"Team ID"
// @Param		Team				body		string 	true 	"Team"
// @Accept json
// @Produce json
// @Success 200 {object} equipes.UpdateEquipe
// @Failure 401,400 {array} utils.ResError
// @Tags Teams
// @Router /teams/{id} [put]
func atualizarEquipe(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
			"Could not update team. Parameters were not passed correctly", 400))
		return
	}
	if res, err := equipe.AtualizarEquipe(id, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}
