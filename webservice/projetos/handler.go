package projetos

import (
	"net/http"

	"gerenciadorDeProjetos/domain/projetos"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary POST a new Project
// @Description POST a new project. For the request to be met, the "nome_projeto", "equipe_id", "descricao_projeto" are required. The status already goes with a predefined value "A Fazer". the "prazo_entrega" is the number of days that the delivery time will be
// @Param		NewProject		body	string		true	"NewProject"
// @Accept json
// @Produce json
// @Success 201 {object} modelApresentacao.ReqProjeto "OK"
// @Failure 400,401 {array} utils.ResError
// @Tags Projects
// @Router /projetos [post]
func NovoProjeto(c *gin.Context) {

	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
			"Could not create. Parameters were not passed correctly", 400))
		return
	}
	if res, err := projetos.NovoProjeto(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError("", "Team does not exist", 400))
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// @Security bearerAuth
// Get Projects
// @Summary Get All Projects
// @Description Get list all project
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqProjetos
// @Failure 404 {string} string "error"
// @Tags Projects
// @Router /projetos [get]
func ListarProjetos(c *gin.Context) {

	if projetos, err := projetos.ListarProjetos(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

// @Security bearerAuth
// Get Project
// @Summary Get Project with specific ID
// @Description GET a project with a specific ID
// @Param        id   path      int  true  "Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqProjetos
// @Failure 400 {array} string "Project does not exist"
// @Failure 404 {string} string "not authorized"
// @Tags Projects
// @Router /projetos/{id} [get]
func ListarProjeto(c *gin.Context) {
	id := c.Param("id")

	if projetos, err := projetos.ListarProjeto(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Project does not exist", 404))
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

// @Security bearerAuth
// Get Tasks of Project
// @Summary Get Tasks of Project with Param ID
// @Description GET all tasks of a project with ID_Projeto specific
// @Param        id   path      int  true  "Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTasksProjeto "OK"
// @Failure 404 {array} string "Project does not exist"
// @Failure 401 {array} string "not authorized"
// @Tags Projects
// @Router /projetos/{id}/tasks [get]
func ListarTasksProjeto(c *gin.Context) {
	id := c.Param("id")

	if projetos, err := projetos.ListarTasksProjeto(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Project does not exist", 404))
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

// @Security bearerAuth
// Get Projects with specific status
// @Summary Get Status of Projects with a specific status with Param Status
// @Description GET all registered projects that have the status passed as a parameter
// @Param        status   path      string  true  "Status"		Enums(A Fazer,Em Andamento,Em Teste,Concluido)
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqStatusProjeto "OK"
// @Failure 404 {string} string "error"
// @Tags Projects
// @Router /projetos/status/{status} [get]
func ListarProjetosComStatus(c *gin.Context) {
	status := c.Param("status")

	if projetos, err := projetos.ListarProjetosComStatus(status); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Status does not exist", 404))
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

// @Security bearerAuth
// @Summary Delete a specific Project
// @Description DELETE a specific project. For the request to be met, the "id_projeto" are required
// @Param        id   		path      	int  	true  	"Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {object} utils.ResOk
// @Failure 404 {array} utils.ResError
// @Failure 401 {array} string "Not Authorized"
// @Tags Projects
// @Router /projetos/{id} [delete]
func DeletarProjeto(c *gin.Context) {
	id := c.Param("id")

	if err := projetos.DeletarProjeto(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Project deleted successfully", 200))
	}
}

// PUT Project
// @Security bearerAuth
// @Summary PUT Project with ID
// @Description PUT a specific project. For the request to be met, the "nome_projeto" and "equipe_id" and "descricao_projeto" are required
// @Param        id   				path      	int  	true  	"Projeto ID"
// @Param		Project				body		string 	true 	"Project"
// @Accept json
// @Produce json
// @Success 200 {object} projetos.ReqAtualizarProjetoData
// @Failure 400,401 {array} utils.ResError
// @Tags Projects
// @Router /projetos/{id} [put]
func AtualizarProjeto(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqAtualizarProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
			"Could not update project. Parameters were not passed correctly", 400))
		return
	}

	if res, err := projetos.AtualizarProjeto(id, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @Security bearerAuth
// @Summary PUT Status of a Project
// @Description PUT Status of a specific project. For the request to be met, the "status" are required
// @Param        id   						path      	int  	true  	"id"
// @Param		Status-Project				body		string 	true 	"Status-Project"
// @Accept json
// @Produce json
// @Success 200 {object} projetos.ReqUpdateStatusProjeto "OK"
// @Failure 400,401 {array} utils.ResError
// @Tags Projects
// @Router /projetos/{id}/status [put]
func AtualizarStatusProjeto(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqAtualizarProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
			"Could not update. Parameters were not passed correctly", 400))
		return
	}

	if res, err := projetos.AtualizarStatusProjeto(id, &req); err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}
