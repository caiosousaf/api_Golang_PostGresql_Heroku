package equipes

import (
	"fmt"
	"net/http"

	"gerenciadorDeProjetos/domain/equipes"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
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

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes")
	if equipes, err := equipe.ListarEquipes(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, equipes)
	}
}

func buscarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar equipe")
	if equipe, err := equipe.BuscarEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func buscarMembrosDeEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar os membros de uma equipe")
	if equipe, err := equipe.BuscarMembrosDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func buscarProjetosDeEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar os Projetos de uma equipe")
	if equipe, err := equipe.BuscarProjetosDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func buscarTasksDeEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar as Tarefas de uma equipe")
	if equipe, err := equipe.BuscarTasksDeEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, utils.KeyError(err.Error(), "Team does not exist", 404))
	} else if len(equipe) == 0 {
		c.JSON(http.StatusNoContent, equipe)
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func deletarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma equipe")

	if err := equipe.DeletarEquipe(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, utils.KeyOk("Team deleted successfully", 200))
	}
}

func atualizarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar equipe")

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