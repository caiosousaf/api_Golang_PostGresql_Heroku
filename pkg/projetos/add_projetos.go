package projetos

import (
	"net/http"
	"strconv"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	Nome_Projeto 		string `json:"nome_projeto"`
	Descricao_Projeto 	string `json:"descricao_projeto"`
	EquipeID 	 		string `json:"equipe_id"`
}

func (h handler) AddProjeto(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	if eqId, err := strconv.Atoi(body.EquipeID); err == nil{
		projeto.Nome_Projeto = body.Nome_Projeto
		projeto.Descricao_Projeto = body.Descricao_Projeto
		projeto.EquipeID = eqId
		projeto.Status = "Em planejamento"
	}

	var check int

	sql := "select count(id_pessoa) from pessoas where equipe_id = ?"

	if result := h.DB.Raw(sql, projeto.EquipeID).Scan(&check); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if(check > 0){
		if result := h.DB.Create(&projeto); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		c.JSON(http.StatusCreated, &projeto)
	} else{
		c.JSON(http.StatusOK, gin.H{"Message":"Projetos só podem ser cadastrados em equipes que possuam membros."})
	}

}