package tasks

import (
	"database/sql"
	"gerenciadorDeProjetos/infra/tasks/postgres"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	modelData "gerenciadorDeProjetos/infra/tasks/model"
)

type repositorio struct {
	Data *postgres.DBTasks
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBTasks{DB: novoDB},
	}
}

func (r *repositorio) NovaTask(req *modelApresentacao.ReqTaskApresent) (*modelApresentacao.ReqTask, error) {
	return r.Data.NovaTask(&modelData.ReqTaskData{Descricao_Task: req.Descricao_Task, PessoaID: req.PessoaID,
	ProjetoID: req.ProjetoID, Prazo: req.Prazo, Prioridade: req.Prioridade,})
}