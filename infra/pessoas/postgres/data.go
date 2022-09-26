package postgres

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/pessoas/model"
	utils "gerenciadorDeProjetos/utils/params"

	sq "github.com/Masterminds/squirrel"
)

type DBPessoas struct {
	DB *sql.DB
}

func (postgres *DBPessoas) NovaPessoa(req *modelData.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	sqlStatement := `INSERT INTO pessoas (nome_pessoa, funcao_pessoa, equipe_id)
					 VALUES ($1, $2, $3) RETURNING *`

	var pessoa = &modelApresentacao.ReqPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.Equipe_ID)
	if err := row.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa, &pessoa.Equipe_ID,
		&pessoa.Data_Contratacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return pessoa, nil
}

func (pg *DBPessoas) ListarPessoas() (res *modelApresentacao.ListarGetPessoa,err error) {
	//sqlStatement := `SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao
	//FROM pessoas as pe INNER JOIN equipes as eq on pe.equipe_id = eq.id_equipe ORDER BY pe.id_pessoa`
	var pessoa = modelApresentacao.ReqGetPessoa{}

	sqlStatement, sqlValues, err := sq.
		Select("pe.*, eq.nome_equipe").
		From("pessoas pe").
		Join("equipes eq ON eq.id_equipe = pe.equipe_id").
		OrderBy("pe.id_pessoa").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}
	rows, err := pg.DB.Query(sqlStatement, sqlValues...)
	if err != nil {
		return nil, err
	}

	res = &modelApresentacao.ListarGetPessoa{
		Pessoas: make([]modelApresentacao.ReqGetPessoa, 0),
	}

	for rows.Next() {
		if err := rows.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
			&pessoa.EquipeID, &pessoa.Nome_Equipe, &pessoa.Data_Contratacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} 
			return nil, err
		}
		res.Pessoas = append(res.Pessoas, pessoa)
	}

	return res, nil
}

func (pg *DBPessoas) ListarPessoa(id string) (res *modelApresentacao.ReqGetPessoa, err error) {

	sqlStatement1, sqlValues, err := sq.
		Select("pe.*, eq.nome_equipe").
		From("pessoas pe").
		Join("equipes eq ON pe.equipe_id = eq.id_equipe").
		Where(sq.Eq{
			"pe.id_pessoa": id,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var pessoa = &modelApresentacao.ReqGetPessoa{}

	row := pg.DB.QueryRow(sqlStatement1, sqlValues...)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
		&pessoa.EquipeID, &pessoa.Data_Contratacao, &pessoa.Nome_Equipe); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return pessoa, nil
}

func (postgres *DBPessoas) ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error) {
	sqlStatement := `	SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, eq.id_equipe, eq.nome_equipe, pr.nome_projeto,tk.id_task, tk.descricao_task,
	tk.projeto_id, tk.status, tk.data_criacao, tk.data_conclusao, tk.prazo_entrega, tk.prioridade 
	FROM pessoas pe 
	INNER JOIN equipes eq ON pe.equipe_id = eq.id_equipe 
	INNER JOIN projetos pr ON pr.equipe_id = eq.id_equipe 
	INNER JOIN tasks tk ON tk.pessoa_id = pe.id_pessoa 
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

	return res, nil
}

func (postgres *DBPessoas) AtualizarPessoa(id string, req *modelData.ReqPessoa) (*modelApresentacao.ReqAtualizarPessoa, error) {
	sqlStatement := `UPDATE pessoas 
					 SET nome_pessoa = $1, funcao_pessoa = $2, equipe_id = $3 
					 WHERE id_pessoa = $4 RETURNING nome_pessoa, funcao_pessoa, equipe_id`
	var pessoa = &modelApresentacao.ReqAtualizarPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.Equipe_ID, id)

	if err := row.Scan(&pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa, &pessoa.Equipe_ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return pessoa, nil
}

func (postgres *DBPessoas) DeletarPessoa(id string) error {
	sqlStatement := `DELETE FROM pessoas WHERE id_pessoa = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func (pg *DBPessoas) ListarPessoasFiltro(params *utils.RequestParams) (res *modelApresentacao.ListarGetPessoa, err error) {
	var (
		ordem, ordenador string
		
	)

	if params.TemFiltro("order") {
		ordem = params.Filters["order"][0]
	}

	if params.TemFiltro("orderBy") {
		ordenador = params.Filters["orderBy"][0]
	}

	sqlStmt, sqlValues, err := sq.
		// Select("pe.*, eq.nome_equipe").
		// From("pessoas pe").
		// Join("equipes eq ON eq.id_equipe = pe.equipe_id").
		// OrderBy(ordenador + " " + ordem).
		// PlaceholderFormat(sq.Dollar).
		// ToSql()

		Select("pe.*, eq.nome_equipe").
		From("pessoas pe").
		Join("equipes eq ON eq.id_equipe = pe.equipe_id").
		Where(sq.ILike{
			ordenador : "%"+ordem+"%",
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := pg.DB.Query(sqlStmt, sqlValues...)
	if err != nil {
		return nil, err
	}

	var pessoa = modelApresentacao.ReqGetPessoa{}

	res = &modelApresentacao.ListarGetPessoa{
		Pessoas: make([]modelApresentacao.ReqGetPessoa, 0),
	}

	for rows.Next() {
		if err := rows.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
			&pessoa.EquipeID, &pessoa.Data_Contratacao, &pessoa.Nome_Equipe); err != nil {
			if err == sql.ErrNoRows {
				return res, nil
			}
			return nil, err
		}
		res.Pessoas = append(res.Pessoas, pessoa)
	}
	return
}