package postgres

import (
	"database/sql"
	"fmt"
	"net/http"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	"github.com/gin-gonic/gin"
)

type DBProjetos struct {
	DB *sql.DB
}

func (postgres *DBProjetos) NovoProjeto(req *modelData.ReqProjeto, c *gin.Context) {
	sqlStatement := `INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES($1, $2 , $3, current_date+$4 );`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Projeto, req.Descricao_Projeto, req.Equipe_ID, req.Prazo)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("Cadastro de novo projeto deu certo")
}