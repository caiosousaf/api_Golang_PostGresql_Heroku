package pessoas

import (
	"net/http"
	"time"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type GetPessoasBody struct{
	ID_Pessoa 		uint 			`json:"id_pessoa"`
	Nome_Pessoa 	string 			`json:"nome_pessoa"`
	Funcao_Pessoa 	string 			`json:"funcao_pessoa"`
	DataContratacao time.Time		`json:"data_contratacao"`
	EquipeID 		*int 			`json:"equipe_id"`
	Equipe 			*models.Equipe 	`json:"equipe"`
	Tasks 			*[]models.Task	`json:"tasks"`
	Favoritar		uint			`json:"favoritar"`
}
func (h handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")

	var pessoas models.Pessoa

	if result := h.DB.Raw("select * from pessoas where id_pessoa = ?", id).Scan(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var eq *int
	var equipe *models.Equipe = nil

	if pessoas.EquipeID == 0{
		eq = nil
	} else{
		eq = &pessoas.EquipeID
		if result := h.DB.First(&equipe, eq); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	}

	var tasks []models.Task
	if result := h.DB.Raw("select * from tasks where pessoa_id = ?", id).Scan(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var data time.Time
	if result := h.DB.Raw("select data_contratacao from pessoas where id_pessoa = ?", id).Scan(&data); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var pessoa GetPessoasBody

	pessoa.ID_Pessoa = pessoas.ID_Pessoa
	pessoa.Nome_Pessoa = pessoas.Nome_Pessoa
	pessoa.Funcao_Pessoa = pessoas.Funcao_Pessoa
	pessoa.DataContratacao = data
	pessoa.EquipeID = eq
	pessoa.Equipe = equipe
	pessoa.Tasks = &tasks
	pessoa.Favoritar = pessoas.Favoritar

	c.JSON(http.StatusOK, &pessoa)
}