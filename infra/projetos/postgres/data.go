package postgres

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	utils "gerenciadorDeProjetos/utils/params"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type DBProjetos struct {
	DB *sql.DB
}

func (postgres *DBProjetos) NovoProjeto(req *modelData.ReqProjeto) (*modelApresentacao.ReqProjetos, error) {
	var t = req.Prazo
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)
	sqlStatement := `INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) 
					 VALUES($1, $2 , $3, $4) RETURNING *`
	var projeto = &modelApresentacao.ReqProjetos{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Projeto, req.Descricao_Projeto, req.Equipe_ID, data_limite)
	if err := row.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto, &projeto.EquipeID, &projeto.Status,
		&projeto.Data_Criacao, &projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return projeto, nil
}

func (postgres *DBProjetos) ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	sqlStatement := `SELECT pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega
					 FROM projetos AS pr 
					 INNER JOIN equipes AS eq ON pr.equipe_id = eq.id_equipe ORDER BY id_projeto`

	var projeto = modelApresentacao.ReqProjetos{}
	var res = []modelApresentacao.ReqProjetos{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
			&projeto.EquipeID, &projeto.Nome_Equipe, &projeto.Status, &projeto.Data_Criacao,
			&projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, projeto)
	}
	return res, nil
}

func (postgres *DBProjetos) ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error) {
	sqlStatement := `SELECT pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega
					 FROM projetos AS pr 
					 INNER JOIN equipes AS eq ON pr.equipe_id = eq.id_equipe
					 WHERE pr.id_projeto = $1`

	var projeto = &modelApresentacao.ReqProjetos{}

	rows := postgres.DB.QueryRow(sqlStatement, id)

	if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
		&projeto.EquipeID, &projeto.Nome_Equipe, &projeto.Status, &projeto.Data_Criacao,
		&projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return projeto, nil
}

func (postgres *DBProjetos) ListarTasksProjeto(id string) ([]modelApresentacao.ReqTasksProjeto, error) {
	sqlStatement := `select tk.*, pr.id_projeto, pr.nome_projeto, eq.nome_equipe,
					 pe.nome_pessoa from 
					 projetos as pr inner join tasks as tk on pr.id_projeto = tk.projeto_id inner join
					 equipes as eq on pr.equipe_id = eq.id_equipe inner join
					 pessoas as pe on pe.id_pessoa = tk.pessoa_id where id_projeto = $1`

	var projeto = modelApresentacao.ReqTasksProjeto{}
	var res = []modelApresentacao.ReqTasksProjeto{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Task, &projeto.Descricao_Task, &projeto.Pessoa_ID, &projeto.Projeto_ID, &projeto.Status,
			&projeto.Prioridade, &projeto.Data_Criacao, &projeto.Data_Conclusao, &projeto.Prazo_Entrega,
			&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Nome_Equipe, &projeto.Nome_Pessoa); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, projeto)
	}

	return res, nil
}

func (postgres *DBProjetos) ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error) {
	sqlStatement := `SELECT * FROM projetos WHERE status = $1`

	var projeto = modelApresentacao.ReqStatusProjeto{}
	var res = []modelApresentacao.ReqStatusProjeto{}

	rows, err := postgres.DB.Query(sqlStatement, status)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto, &projeto.EquipeID,
			&projeto.Status, &projeto.Data_Criacao,
			&projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, projeto)
	}

	return res, nil
}

func (postgres *DBProjetos) DeletarProjeto(id string) error {
	sqlStatement := `DELETE FROM projetos where id_projeto = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func (postgres *DBProjetos) AtualizarProjeto(id string, req *modelData.ReqAtualizarProjetoData) (*modelApresentacao.ReqAtualizarProjeto, error) {
	sqlStatement := `UPDATE projetos
					 SET nome_projeto = $1, equipe_id = $2, descricao_projeto = $3
					 WHERE id_projeto = $4 RETURNING *`

	var projeto = &modelApresentacao.ReqAtualizarProjeto{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Projeto, req.Equipe_ID, req.Descricao_Projeto, id)

	if err := row.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto,  &projeto.Descricao_Projeto, &projeto.EquipeID, &projeto.Status,
		&projeto.Data_Criacao, &projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return projeto, nil
}

func (postgres *DBProjetos) AtualizarStatusProjeto(id string, req *modelData.ReqUpdateStatusProjeto) (*modelApresentacao.ReqAtualizarProjeto, error) {
	sqlStatement := `UPDATE projetos
					 SET status = $1 
					 WHERE id_projeto = $2 RETURNING *`

	var projeto = &modelApresentacao.ReqAtualizarProjeto{}

	row := postgres.DB.QueryRow(sqlStatement, req.Status, id)

	if err := row.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto, &projeto.EquipeID, &projeto.Status, 
		&projeto.Data_Criacao, &projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	sqlStatementStatus := `update projetos set data_conclusao = current_date where status = 'Concluido' and id_projeto = $1`
	err := postgres.DB.QueryRow(sqlStatementStatus, id)
	if err != nil {
		return projeto, nil
	}

	return projeto, nil
}

func (pg *DBProjetos) ListarProjetosFiltro(params *utils.RequestParams) (res []modelApresentacao.ReqProjetos, err error) {
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
		Select(`pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega`).
		From("projetos AS pr").
		Join("equipes AS eq ON pr.equipe_id = eq.id_equipe").
		Where(sq.ILike{
			column : "%"+value+"%",
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	} 
	if !params.TemFiltro("value") && !params.TemFiltro("column") && !params.TemFiltro("order") && !params.TemFiltro("orderBy"){
		sqlStmt, sqlValues, err = sq.
		Select(`pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega`).
		From("projetos AS pr").
		Join("equipes AS eq ON pr.equipe_id = eq.id_equipe").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	}

	if params.TemFiltro("order") && params.TemFiltro("orderBy")  {
		sqlStmt, sqlValues, err = sq.
		Select(`pr.id_projeto, pr.nome_projeto,pr.descricao_projeto, pr.equipe_id, eq.nome_equipe, pr.status, 
					 pr.data_criacao, pr.data_conclusao, pr.prazo_entrega`).
		From("projetos AS pr").
		Join("equipes AS eq ON pr.equipe_id = eq.id_equipe").
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

	var projeto = modelApresentacao.ReqProjetos{}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
			&projeto.EquipeID, &projeto.Nome_Equipe, &projeto.Status, &projeto.Data_Criacao,
			&projeto.Data_Conclusao, &projeto.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}   
		res = append(res, projeto)
	}
	return res, nil
}