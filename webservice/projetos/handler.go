package projetos

import (
	"database/sql"
	"fmt"

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

	projetos.NovoProjeto(&req, c)
	c.JSON(201, gin.H{"OK": "Projeto cadastrado com sucesso"})
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