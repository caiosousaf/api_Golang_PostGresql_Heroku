package pessoas

import (
	"database/sql"
	"fmt"
	"gerenciadorDeProjetos/domain/pessoas"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NovaPessoa(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova pessoa")
	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "error": err.Error(),
		})
		return
	}

	pessoas.NovaPessoa(&req, c)
	c.JSON(http.StatusCreated, gin.H{"OK": "Pessoa Cadastrada com sucesso"})
}

func ListarPessoas(c *gin.Context) {
	fmt.Println("Tentando Listar todas as pessoas")
	if pessoas, err := pessoas.ListarPessoas(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": "Nenhum registro encontrado", "err": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar uma pessoa com id especifico")
	if pessoas, err := pessoas.ListarPessoa(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": "Nenhum registro encontrado", "err": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarTarefasPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar tarefas de uma pessoa com id especifico")
	if pessoas, err := pessoas.ListarTarefasPessoa(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": "Nenhum registro encontrado", "err": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func AtualizarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar os dados de uma pessoa")

	req := modelApresentacao.ReqAtualizarPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
		return
	}

	if res, err := pessoas.AtualizarPessoa(id, &req); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": "Nenhum registro encontrado", "err": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func DeletarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma pessoa")
	if _, err := pessoas.ListarPessoa(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"message": "Nenhum registro encontrado", "err": err.Error(),
			})
			return
		} else {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
	} else {
		pessoas.DeletarPessoa(id)
		c.JSON(http.StatusOK, gin.H{"Message": "Pessoa deletada com sucesso"})
	}
}
