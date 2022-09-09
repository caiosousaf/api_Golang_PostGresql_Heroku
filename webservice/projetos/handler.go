package projetos

import (
	"fmt"
	"net/http"

	"gerenciadorDeProjetos/domain/projetos"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"

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

func ListarProjetos(c *gin.Context) {
	fmt.Println("Tentando listar todos os projetos")
	projetos, err := projetos.ListarProjetos()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, projetos)
	}
}

func ListarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar um projeto")
	projetos, err := projetos.ListarProjeto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "Message": "Projeto não existe"})
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