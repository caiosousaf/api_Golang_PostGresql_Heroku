package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
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

func (postgres *DBPessoas) ListarPessoas() ([]modelApresentacao.ReqGetPessoa, error) {
	sqlStatement := `SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao
	FROM pessoas as pe INNER JOIN equipes as eq on pe.equipe_id = eq.id_equipe ORDER BY pe.id_pessoa`
	var pessoa = modelApresentacao.ReqGetPessoa{}
	var res = []modelApresentacao.ReqGetPessoa{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
			&pessoa.EquipeID, &pessoa.Nome_Equipe, &pessoa.Data_Contratacao); err != nil {
			return nil, err
		}
		res = append(res, pessoa)
	}
	fmt.Println("Listagem de todas as pessoas deu certo!!")
	return res, nil
}

func (postgres *DBPessoas) ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error) {
	sqlStatement := `select pe.*, eq.nome_equipe
					 from pessoas as pe 
					 inner join equipes as eq on pe.equipe_id = eq.id_equipe 
					 where id_pessoa = $1`
	var pessoa = &modelApresentacao.ReqGetPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
		 &pessoa.EquipeID, &pessoa.Data_Contratacao, &pessoa.Nome_Equipe); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	fmt.Println("Buscar de uma pessoa deu certo!!")
	return pessoa, nil
}