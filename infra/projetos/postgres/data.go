package postgres

import (
	"database/sql"
	"fmt"
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