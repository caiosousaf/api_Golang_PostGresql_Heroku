package equipes

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)
type GetProjectsMembersTasks struct {
	ID_Equipe   uint             			`json:"id_equipe"`
	Nome_Equipe string           			`json:"nome_equipe"`
	Pessoas     []models.Pessoa  			`json:"pessoas"`
	Projetos    []models.Projeto 			`json:"projetos"`
	Tasks		[]models.TasksbyTeam		`json:"tasks"`		
}

// @Summary Get Specific Team
// @Description Returns a team, all their members, all projects and all tasks they are assigned to
// @Param		id		path	int		true	"id_Team"
// @Accept json
// @Produce json
// @Success 200 {array} GetProjectsMembersTasks
// @Failure 400,404 {string} string "error"
// @Tags Teams
// @Router /equipes/{id} [get]
func (h handler) GetTeam(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe
	var me GetProjectsMembersTasks

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var pessoas []models.Pessoa
	var projetos []models.Projeto
	var tasks	[]models.TasksbyTeam
	if result := h.DB.Where("equipe_id = ?", equipe.ID_Equipe).Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Where("equipe_id = ?", equipe.ID_Equipe).Find(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw(`select tk.id_task, pe.nome_pessoa, tk.pessoa_id, tk.descricao_task, tk.projeto_id, tk.status,
	tk.data_criacao, tk.prazo_entrega, tk.data_conclusao
	from tasks tk 
	inner join pessoas pe on pe.id_pessoa = tk.pessoa_id 
	inner join equipes eq on eq.id_equipe = pe.equipe_id	
	where eq.id_equipe = ?`, equipe.ID_Equipe).Scan(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	me.ID_Equipe = equipe.ID_Equipe
	me.Nome_Equipe = equipe.Nome_Equipe
	me.Pessoas = pessoas
	me.Projetos = projetos
	me.Tasks = tasks

	c.JSON(http.StatusOK, &me)
}