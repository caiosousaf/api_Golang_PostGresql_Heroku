package pessoas

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/models"
)

type AddPessoaRequestBody struct {
    Nome_Pessoa		string `gorm:"type: varchar(30) not null" json:"nome_pessoa"`
	Funcao_Pessoa	string `gorm:"type: varchar(15) not null" json:"funcao_pessoa"`
	Equipe			models.Equipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}

func (h handler) AddPerson(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.Equipe = body.Equipe

	if result := h.DB.Create(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pessoa)
}