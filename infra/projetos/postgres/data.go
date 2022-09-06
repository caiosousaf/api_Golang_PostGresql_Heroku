package postgres

import (
	"database/sql"
	"fmt"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DBProjetos struct {
	DB *sql.DB
}

func (postgres *DBProjetos) NovoProjeto(req *modelData.ReqProjeto, c *gin.Context) {
	var t = req.Prazo
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)
	sqlStatement := `INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES($1, $2 , $3, $4);`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Projeto, req.Descricao_Projeto, req.Equipe_ID, data_limite)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("Cadastro de novo projeto deu certo")
}

func (postgres *DBProjetos) ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	sqlStatement := `SELECT pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega
					 FROM projetos AS pr 
					 INNER JOIN equipes AS eq ON pr.equipe_id = eq.id_equipe`

	var projeto = modelApresentacao.ReqProjetos{}
	var res = []modelApresentacao.ReqProjetos{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
			&projeto.EquipeID, &projeto.Nome_Equipe,&projeto.Status, &projeto.Data_Criacao, 
			&projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
			return nil, err
		}
		res = append(res, projeto)
	}
	fmt.Println("Listagem de todas os projetos deu certo!!")
	return res, nil
}