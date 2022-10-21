package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/equipes/model"
	utils "gerenciadorDeProjetos/utils/params"

	sq "github.com/Masterminds/squirrel"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT) RETURNING *`

	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		return nil, err
	}
	return equipe, nil
}

func (postgres *DBEquipes) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe`
	var res = []modelApresentacao.ReqEquipe{}
	var equipe = modelApresentacao.ReqEquipe{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}

		res = append(res, equipe)
	}
	return res, nil
}

func (postgres *DBEquipes) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes WHERE id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	return equipe, nil
}

func (postgres *DBEquipes) BuscarMembrosDeEquipe(id string) ([]modelPessoa.ReqMembros, error) {
	sqlStatement := `select id_pessoa, nome_pessoa, funcao_pessoa, equipe_id, data_contratacao from pessoas WHERE equipe_id = $1`
	var res = []modelPessoa.ReqMembros{}
	var equipe = modelPessoa.ReqMembros{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.ID_Pessoa, &equipe.Nome_Pessoa, &equipe.Funcao_Pessoa, &equipe.Equipe_ID, &equipe.Data_Contratacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	return res, nil
}

func (postgres *DBEquipes) BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error) {
	sqlStatement := `select eq.nome_equipe, pr.id_projeto, pr.nome_projeto, pr.status, pr.descricao_projeto, pr.data_criacao, pr.data_conclusao, pr.prazo_entrega 
	from equipes as eq 
	inner join projetos as pr on eq.id_equipe = pr.equipe_id where eq.id_equipe = $1`
	var res = []modelApresentacao.ReqEquipeProjetos{}
	var equipe = modelApresentacao.ReqEquipeProjetos{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.Nome_Equipe, &equipe.ID_Projeto, &equipe.Nome_Projeto, &equipe.Status, &equipe.Descricao_Projeto,
			&equipe.Data_Criacao, &equipe.Data_Conclusao, &equipe.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	return res, nil
}

func (postgres *DBEquipes) BuscarTasksDeEquipe(id string) ([]modelApresentacao.ReqTasksbyTeam, error) {
	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, tk.status,
	tk.data_criacao, tk.prazo_entrega, tk.data_conclusao, tk.prioridade
					 from tasks tk 
					 inner join pessoas pe on pe.id_pessoa = tk.pessoa_id 
					 inner join equipes eq on eq.id_equipe = pe.equipe_id	
					 where eq.id_equipe = $1`
	var res = []modelApresentacao.ReqTasksbyTeam{}
	var equipe = modelApresentacao.ReqTasksbyTeam{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Task, &equipe.Descricao_Task, &equipe.Pessoa_ID, &equipe.Nome_Pessoa, &equipe.Projeto_ID,
			&equipe.Status, &equipe.Data_Criacao, &equipe.Prazo_Entrega, &equipe.Data_Conclusao, &equipe.Prioridade); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	return res, nil
}

func (postgres *DBEquipes) DeletarEquipe(id string) error {
	sqlStatement := `DELETE FROM equipes WHERE id_equipe = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (postgres *DBEquipes) AtualizarEquipe(id string, req *modelData.UpdateEquipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement, sqlValues, err := sq.
	Update("equipes").
	Set("nome_equipe", req.Nome_Equipe).
	Where(sq.Eq{"id_equipe": id}).
	PlaceholderFormat(sq.Dollar).
	Suffix("RETURNING *").
	ToSql()
	fmt.Println(sqlStatement)
	if err != nil {
		return nil, err
	}
	// sqlStatement := `UPDATE equipes SET nome_equipe = $1 
	// WHERE id_equipe = $2 RETURNING *`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, sqlValues...)

	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return equipe, nil
}

func (pg *DBEquipes) ListarEquipesFiltro(params *utils.RequestParams) (res []modelApresentacao.ReqEquipe, err error) {
	var (
		ordem, ordenador string
		column, value string
	)

	if params.TemFiltro("value") {
		value = params.Filters["value"][0]
	}

	if params.TemFiltro("column") {
		column = params.Filters["column"][0]
	}

	if params.TemFiltro("orderBy") {
		ordenador = params.Filters["orderBy"][0]
	}

	if params.TemFiltro("order") {
		ordem = params.Filters["order"][0]
	}


	var sqlStmt string
	var sqlValues []interface{}

	if params.TemFiltro("value") && params.TemFiltro("column")  {
		sqlStmt, sqlValues, err = sq.
		Select("*").
		From("equipes").
		Where(sq.ILike{
			column : "%"+value+"%",
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	} 
	if !params.TemFiltro("value") && !params.TemFiltro("column") && !params.TemFiltro("order") && !params.TemFiltro("orderBy"){
		sqlStmt, sqlValues, err = sq.
		Select("*").
		From("equipes").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	}

	if params.TemFiltro("order") && params.TemFiltro("orderBy")  {
		sqlStmt, sqlValues, err = sq.
		Select("*").
		From("equipes").
		OrderBy(ordenador + " " + ordem).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	}
	
	if err != nil {
		return nil, err
	}

	rows, err := pg.DB.Query(sqlStmt, sqlValues...)
	if err != nil {
		return nil, err
	}

	var equipe = modelApresentacao.ReqEquipe{}

	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}

		res = append(res, equipe)
	}
	return res, nil
}