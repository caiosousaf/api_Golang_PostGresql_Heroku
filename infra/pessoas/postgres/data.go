package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

	modelData "gerenciadorDeProjetos/infra/pessoas/model"

	"github.com/gin-gonic/gin"
)

type DBPessoas struct {
	DB *sql.DB
}

func (postgres *DBPessoas) NovaPessoa(req *modelData.ReqPessoa, c *gin.Context) {
	sqlStatement := `INSERT INTO pessoas (nome_pessoa, funcao_pessoa, equipe_id)
					 VALUES ($1, $2, $3)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.Equipe_ID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("Cadastro de nova pessoa deu certo")
}