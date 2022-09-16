package pessoas

import (
	"fmt"
	"gerenciadorDeProjetos/domain/pessoas"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"net/http"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

func NovaPessoa(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova pessoa")
	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(),
		"Could not create. Parameters were not passed correctly", 400))
		return
	}
	if res, err := pessoas.NovaPessoa(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(), "Team does not exist", 400))
	} else {
		c.JSON(http.StatusCreated, res)
	}

}

func ListarPessoas(c *gin.Context) {
	fmt.Println("Tentando Listar todas as pessoas")
	if pessoas, err := pessoas.ListarPessoas(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar uma pessoa com id especifico")
	if pessoas, err := pessoas.ListarPessoa(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Person does not exist", 404))
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

func ListarTarefasPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar tarefas de uma pessoa com id especifico")
	if pessoas, err := pessoas.ListarTarefasPessoa(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Person does not exist", 404))
	} else if len(pessoas) == 0 {
		c.JSON(http.StatusNoContent, pessoas)
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

	if res, err := pessoas.AtualizarPessoa(id, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func DeletarPessoa(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma pessoa")

	if err := pessoas.DeletarPessoa(id); err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Person deleted successfully", 200))
	}
}
