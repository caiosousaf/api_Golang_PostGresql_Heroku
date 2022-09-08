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

func (r *repositorio) ListarTasks() ([]modelApresentacao.ReqTasks, error) {
	return r.Data.ListarTasks()
}

func (r *repositorio) ListarTask(id string) (*modelApresentacao.ReqTasks, error) {
	return r.Data.ListarTask(id)
}

func (r *repositorio) ListarStatusTasks(status string) ([]modelApresentacao.ReqTasks, error) {
	return r.Data.ListarStatusTasks(status)
}

func (r *repositorio) AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error){
	return r.Data.AtualizarTask(id, &modelData.ReqUpdateTaskData{Descricao_Task: req.Descricao_Task, PessoaID: req.PessoaID, ProjetoID: req.ProjetoID, Prioridade: req.Prioridade})
} 

func (r *repositorio) AtualizarStatusTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error){
	return r.Data.AtualizarStatusTask(id, &modelData.ReqUpdateStatusTask{Status: req.Status})
} 

func (r *repositorio) DeletarTask(id string) error {
	return r.Data.DeletarTask(id)
}
