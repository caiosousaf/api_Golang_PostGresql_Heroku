package projetos

import (
	"fmt"
	"net/http"

	"gerenciadorDeProjetos/domain/projetos"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

func NovoProjeto(c *gin.Context) {
	fmt.Println("Tentando cadastrar um novo projeto")
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
	fmt.Println("Tentando listar todos os projetos")
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
	fmt.Println("Tentando listar um projeto")
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
	fmt.Println("Tentando listar tarefas de um projeto")
	if projetos, err := projetos.ListarTasksProjeto(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Project does not exist", 404))
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

func ListarProjetosComStatus(c *gin.Context) {
	status := c.Param("status")
	fmt.Println("Tentando listar todos os projetos com um status especifico")
	if projetos, err := projetos.ListarProjetosComStatus(status); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Status does not exist", 404))
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

func DeletarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar um projeto")
	if err := projetos.DeletarProjeto(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Project deleted successfully", 200))
	}
}

func AtualizarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar um projeto")

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

func AtualizarStatusProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar status de um projeto")

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