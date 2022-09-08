package projetos

import (
	"database/sql"
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
			"message": "Could not create. Parameters were not passed correctly " , "error": err.Error(),
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
	if projetos, err := projetos.ListarProjetos(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(200, projetos)
	}
}

func ListarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar um projeto")
	if projetos, err := projetos.ListarProjeto(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(200, projetos)
	}
}

func ListarProjetosComStatus(c *gin.Context) {
	status := c.Param("status")
	fmt.Println("Tentando listar todos os projetos com um status especifico")
	if projetos, err := projetos.ListarProjetosComStatus(status); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(200, projetos)
	}
}

func DeletarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar um projeto")
	if _, err := projetos.ListarProjeto(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum projeto encontrado para ser deletado encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		projetos.DeletarProjeto(id)
		c.JSON(200, gin.H{"OK": "Projeto deletado com sucesso"})
	}
}

func AtualizarProjeto(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar um projeto")

	req := modelApresentacao.ReqAtualizarProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
		return
	}

	if res, err := projetos.AtualizarProjeto(id, &req); err != nil {
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}