package pessoas

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddPessoaRequestBody struct {
	Nome_Pessoa		string  `json:"nome_pessoa"`
	Funcao_Pessoa	string  `json:"funcao_pessoa"`
	EquipeID		int  	`json:"equipe_id"`
}

func (h handler) AddPessoa(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	if body.EquipeID != 0{
			pessoa.Nome_Pessoa = body.Nome_Pessoa
			pessoa.Funcao_Pessoa = body.Funcao_Pessoa
			pessoa.EquipeID = body.EquipeID
		if result := h.DB.Create(&pessoa); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		
	} else {
		pessoa.Nome_Pessoa = body.Nome_Pessoa
		pessoa.Funcao_Pessoa = body.Funcao_Pessoa
		if result := h.DB.Raw(`INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) 
		VALUES(?,?,NULL)`, pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa).Scan(&pessoa).Last(&pessoa); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	}


	c.JSON(http.StatusCreated, &pessoa)
}