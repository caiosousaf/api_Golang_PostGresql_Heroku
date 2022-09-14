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
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "error": err.Error(),
		})
		return
	}

	if res, err := projetos.NovoProjeto(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	projetos, err := projetos.ListarProjetos()
	if err != nil {
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
	projetos, err := projetos.ListarProjeto(id)
	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "Message": "Project does not exist"})
		c.JSON(http.StatusNotFound, 
			utils.KeyError(err.Error(), "Casa", 404))
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
	projetos, err := projetos.ListarTasksProjeto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "Message": "Project does not exist"})
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}


func ListarProjetosComStatus(c *gin.Context) {
	status := c.Param("status")
	fmt.Println("Tentando listar todos os projetos com um status especifico")
	projetos, err := projetos.ListarProjetosComStatus(status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "Message": "Status não existe"})
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

func DeletarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar um projeto")
	err := projetos.DeletarProjeto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Projeto deletado com sucesso"})
	}
}

func AtualizarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar um projeto")

	req := modelApresentacao.ReqAtualizarProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
		return
	}

	res, err := projetos.AtualizarProjeto(id, &req)
	if err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
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