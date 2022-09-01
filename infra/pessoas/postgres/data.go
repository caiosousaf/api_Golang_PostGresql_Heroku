package postgres

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type DBPessoas struct {
	DB *sql.DB
}

func (postgres *DBPessoas) NovaPessoa(req *modelData.Pessoa, c *gin.Context) {
	sqlStatement := `INSERT INTO pessoas (nome_pessoa)`
}