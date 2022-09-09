package pessoas

import (
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "error": err.Error(),
		})
		return
	}
	if res, err := pessoas.NovaPessoa(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Equipe inexistente"})

	} else {
		c.JSON(http.StatusCreated, res)
	}

}

func ListarPessoas(c *gin.Context) {
	fmt.Println("Tentando Listar todas as pessoas")
	pessoas, err := pessoas.ListarPessoas()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar uma pessoa com id especifico")
	pessoas, err := pessoas.ListarPessoa(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarTarefasPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar tarefas de uma pessoa com id especifico")
	pessoas, err := pessoas.ListarTarefasPessoa(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func AtualizarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar os dados de uma pessoa")

	req := modelApresentacao.ReqAtualizarPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
		return
	}

	res, err := pessoas.AtualizarPessoa(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func DeletarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma pessoa")

	err := pessoas.DeletarPessoa(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Pessoa deletada com sucesso"})
	}
}
