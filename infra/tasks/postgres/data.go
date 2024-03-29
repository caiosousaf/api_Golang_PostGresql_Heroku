package postgres

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	modelData "gerenciadorDeProjetos/infra/tasks/model"
	utils "gerenciadorDeProjetos/utils/params"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type DBTasks struct {
	DB *sql.DB
}

func (postgres *DBTasks) NovaTask(req *modelData.ReqTaskData) (*modelApresentacao.ReqTask, error) {
	var t = req.Prazo
	var data_atual = time.Now()
	data_limite := data_atual.AddDate(0, 0, t)
	sqlStatement := `INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) 
					 VALUES($1, $2, $3, $4, $5) RETURNING *`
	var task = &modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, req.Descricao_Task, req.PessoaID, req.ProjetoID,
		req.Prioridade, data_limite)
	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.ProjetoID, &task.Status, &task.Prioridade,
		 &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return task, nil
}

func (postgres *DBTasks) ListarTasks() ([]modelApresentacao.ReqTasks, error) {

	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade 
						from tasks as tk 
						inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa 
						inner join projetos as pr on tk.projeto_id = pr.id_projeto 
						order by id_task`

	var task = modelApresentacao.ReqTasks{}
	var res = []modelApresentacao.ReqTasks{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.Nome_Pessoa, &task.ProjetoID, &task.Nome_Projeto, &task.Status, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, task)
	}

	return res, nil
}

func (postgres *DBTasks) ListarTask(id string) (*modelApresentacao.ReqTasks, error) {
	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade 
						from tasks as tk 
						inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa 
						inner join projetos as pr on tk.projeto_id = pr.id_projeto 
						WHERE id_task = $1`
	var task = &modelApresentacao.ReqTasks{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.Nome_Pessoa, &task.ProjetoID, &task.Nome_Projeto, &task.Status, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return task, nil
}

func (postgres *DBTasks) ListarStatusTasks(status string) ([]modelApresentacao.ReqTasks, error) {
	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade 
						from tasks as tk 
						inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa 
						inner join projetos as pr on tk.projeto_id = pr.id_projeto 
						WHERE tk.status = $1
						order by id_task`
	var task = &modelApresentacao.ReqTasks{}
	var res = []modelApresentacao.ReqTasks{}

	row, err := postgres.DB.Query(sqlStatement, status)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.Nome_Pessoa, &task.ProjetoID, &task.Nome_Projeto, &task.Status, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, *task)
	}

	return res, nil
}

func (postgres *DBTasks) AtualizarTask(id string, req *modelData.ReqUpdateTaskData) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `UPDATE tasks
					SET descricao_task = $1, pessoa_id = $2,
					projeto_id = $3, prioridade = $4 
					WHERE id_task = $5 RETURNING *`
	var task = &modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, req.Descricao_Task, req.PessoaID, req.ProjetoID, req.Prioridade, id)

	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.ProjetoID, &task.Status, &task.Prioridade, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega ); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return task, nil
}

func (postgres *DBTasks) AtualizarStatusTask(id string, req *modelData.ReqUpdateStatusTask) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `UPDATE tasks SET status = $1 WHERE id_task = $2 RETURNING *;`

	var task = &modelApresentacao.ReqTask{}

	sqlStatementStatus := `update tasks set data_conclusao = current_date where status = 'Concluido' and id_task = $1`
	_, err := postgres.DB.Query(sqlStatementStatus, id)
	if err != nil {
		return nil, err
	}
	row := postgres.DB.QueryRow(sqlStatement, req.Status, id)
	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.ProjetoID, &task.Status,&task.Prioridade, &task.Data_Criacao, 
					   &task.Data_Conclusao, &task.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return task, nil
}

func (postgres *DBTasks) DeletarTask(id string) error {
	sqlStatement := `DELETE FROM tasks WHERE id_task = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func (pg *DBTasks) ListarTasksFiltro(params *utils.RequestParams) (res []modelApresentacao.ReqTasks, err error) {
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
		Select(`tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao,
				 tk.data_conclusao,tk.prazo_entrega ,tk.prioridade`).
		From("tasks as tk ").
		Join("pessoas as pe on tk.pessoa_id = pe.id_pessoa").
		Join("projetos as pr on tk.projeto_id = pr.id_projeto").
		Where(sq.ILike{
			column : "%"+value+"%",
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	} 
	if !params.TemFiltro("value") && !params.TemFiltro("column") && !params.TemFiltro("order") && !params.TemFiltro("orderBy"){
		sqlStmt, sqlValues, err = sq.
		Select(`tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao,
				 tk.data_conclusao,tk.prazo_entrega ,tk.prioridade`).
		From("tasks as tk ").
		Join("pessoas as pe on tk.pessoa_id = pe.id_pessoa").
		Join("projetos as pr on tk.projeto_id = pr.id_projeto").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	}

	if params.TemFiltro("order") && params.TemFiltro("orderBy")  {
		sqlStmt, sqlValues, err = sq.
		Select(`tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao,
				 tk.data_conclusao,tk.prazo_entrega ,tk.prioridade`).
		From("tasks as tk ").
		Join("pessoas as pe on tk.pessoa_id = pe.id_pessoa").
		Join("projetos as pr on tk.projeto_id = pr.id_projeto").
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

	var task = modelApresentacao.ReqTasks{}

	for rows.Next() {
		if err := rows.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.Nome_Pessoa, &task.ProjetoID, &task.Nome_Projeto, &task.Status, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, task)
	}

	return res, nil
}