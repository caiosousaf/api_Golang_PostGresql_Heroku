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