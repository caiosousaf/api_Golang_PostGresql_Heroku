package equipes

import (
	"database/sql"
	"fmt"
	"net/http"

	"gerenciadorDeProjetos/domain/equipes"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error(),
		})
		return
	}

	if res, err := equipe.NovaEquipe(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	} else{
		c.JSON(http.StatusCreated, res)
	}
}

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes") 
	if equipes, err := equipe.ListarEquipes(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipes)
	}
}

func buscarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar equipe")
	if equipe, err := equipe.BuscarEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":"" + err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func buscarMembrosDeEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar os membros de uma equipe")
	if equipe, err := equipe.BuscarMembrosDeEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum equipe encontrada", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":"" + err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func buscarProjetosDeEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar os Projetos de uma equipe")
	if equipe, err := equipe.BuscarProjetosDeEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum equipe encontrada", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":"" + err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}

func deletarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar uma equipe")
	if _, err := equipe.BuscarEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhuma equipe encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"error":err.Error()})
		}
	} else {
		equipe.DeletarEquipe(id)
		c.JSON(http.StatusOK, gin.H{"Message": "Equipe deletada com sucesso"})
	}
}

func atualizarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar equipe")

	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not update. Parameters were not passed correctly.", "err":err.Error(),
		})
		return
	}

	if res, err := equipe.AtualizarEquipe(id, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}