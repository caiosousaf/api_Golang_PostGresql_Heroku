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

func (postgres *DBPessoas) ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error) {
	sqlStatement := `	SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, eq.id_equipe, eq.nome_equipe, pr.nome_projeto,tk.id_task, tk.descricao_task,
						tk.projeto_id, tk.status, tk.data_criacao, tk.data_conclusao, tk.prazo_entrega, tk.prioridade 
						FROM pessoas AS pe 
						INNER JOIN equipes AS eq ON pe.equipe_id = eq.id_equipe 
						INNER JOIN projetos AS pr ON pr.equipe_id = eq.id_equipe 
						INNER JOIN tasks as tk ON tk.projeto_id = pr.id_projeto AND tk.pessoa_id = pe.id_pessoa 
						WHERE pe.id_pessoa = $1`
	
	var pessoa = modelApresentacao.ReqTarefaPessoa{}
	var res = []modelApresentacao.ReqTarefaPessoa{}

	row, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
	if err := row.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa, &pessoa.ID_Equipe, &pessoa.Nome_Equipe, 
					   &pessoa.Nome_Projeto, &pessoa.ID_Task, &pessoa.Descricao_Task, &pessoa.Projeto_ID, &pessoa.Status,
					   &pessoa.Data_Criacao, &pessoa.Data_Conclusao, &pessoa.Prazo_Entrega, &pessoa.Prioridade); err != nil {
						if err == sql.ErrNoRows {
							return nil, err
						} else {
							return nil, err
						}
	}
			res = append(res, pessoa)
	}
	fmt.Println("Busca de tarefas de uma pessoa deu certo!!")
	return res, nil
}

func (postgres *DBPessoas) AtualizarPessoa(id string, req *modelData.ReqPessoa) (*modelApresentacao.ReqAtualizarPessoa, error ){
	sqlStatement := `UPDATE pessoas 
					 SET nome_pessoa = $1, funcao_pessoa = $2, equipe_id = $3 
					 WHERE id_pessoa = $4 RETURNING nome_pessoa, funcao_pessoa, equipe_id`
	var pessoa = &modelApresentacao.ReqAtualizarPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.Equipe_ID, id)

	if err := row.Scan(&pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa, &pessoa.Equipe_ID); err != nil {
		return nil, err
	}
	fmt.Println("Atualizar pessoa deu certo")
	return pessoa, nil
}

func (postgres *DBPessoas) DeletarPessoa(id string) error{
	sqlStatement := `DELETE FROM pessoas WHERE id_pessoa = $1`
	
	_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}
		fmt.Println("Tudo certo em deletar um projeto!!")
		return nil
}