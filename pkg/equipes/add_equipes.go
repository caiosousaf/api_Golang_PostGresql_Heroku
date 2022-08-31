//Arquivo com funções para adicionar equipes
package equipes

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddEquipeRequestBody struct {
	Nome_Equipe		string 			`json:"nome_equipe"`
}

func (h handler) AddEquipe(c *gin.Context) {
	//Variável do tipo AddEquipeRequestBody que vai receber as informações passadas através do JSON
	body := AddEquipeRequestBody{}

	//Verificando o corpo de requisição
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//Variável do tipo equipe, que foi definido no pacote models, que vai criar o registro no banco de dados
	var equipe models.Equipe

	//Transferência de informação que o body recebeu para o elemento equipe
	equipe.Nome_Equipe = body.Nome_Equipe

	//Verificação de erros e criação do registro no banco de dados
	if result := h.DB.Create(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//Se retornar o status de criação correto, exibe o registro gravado no banco de dados
	c.JSON(http.StatusCreated, &equipe)
}