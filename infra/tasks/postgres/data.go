package postgres

import (
	"database/sql"
	"fmt"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	modelData "gerenciadorDeProjetos/infra/tasks/model"
	"time"
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
	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.ProjetoID, &task.Status, &task.Data_Criacao,
		&task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
		return nil, err
	}
	fmt.Println("Cadastro de nova task deu certo")
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
			return nil, err
		}
		res = append(res, task)
	}
	fmt.Println("Listagem de todas as tasks deu certo!!")
	return res, nil
}

func (postgres *DBTasks) ListarTask(id string) (*modelApresentacao.ReqTasks, error) {
	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status, tk.data_criacao, tk.data_conclusao,tk.prazo_entrega ,tk.prioridade 
						from tasks as tk 
						inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa 
						inner join projetos as pr on tk.projeto_id = pr.id_projeto 
						WHERE id_task = $1
						order by id_task`
	var task = &modelApresentacao.ReqTasks{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.Nome_Pessoa, &task.ProjetoID, &task.Nome_Projeto, &task.Status, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega, &task.Prioridade); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	fmt.Println("Buscar uma task deu certo!!")
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
	fmt.Println("Busca dos status das tasks deu certo!!")
	return res, nil
}

func (postgres *DBTasks) AtualizarTask(id string, req *modelData.ReqUpdateTaskData) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `UPDATE tasks
					SET descricao_task = $1, pessoa_id = $2,
					projeto_id = $3, prioridade = $4 
					WHERE id_task = $5 RETURNING *`
	var task = &modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, req.Descricao_Task, req.PessoaID, req.ProjetoID, req.Prioridade, id)

	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, &task.ProjetoID, &task.Status, &task.Prioridade, &task.Data_Criacao, &task.Data_Conclusao, &task.Prazo_Entrega); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	fmt.Println("Atualizar uma task deu certo")
	return task, nil
}

func (postgres *DBTasks) DeletarTask(id string) error {
	sqlStatement := `DELETE FROM tasks WHERE id_task = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	fmt.Println("Tudo certo em deletar uma task!!")
	return nil
}