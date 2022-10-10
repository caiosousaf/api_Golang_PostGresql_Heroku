package tasks

import (
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	utils "gerenciadorDeProjetos/utils/params"
)
type ITask interface {
	NovaTask(req *modelApresentacao.ReqTaskApresent) (*modelApresentacao.ReqTask, error)
	ListarTasks() ([]modelApresentacao.ReqTasks, error)
	ListarTask(id string) (*modelApresentacao.ReqTasks, error)
	ListarStatusTasks(status string) ([]modelApresentacao.ReqTasks, error)
	AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error)
	AtualizarStatusTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error)
	DeletarTask(id string) error
	ListarTasksFiltro(params *utils.RequestParams) ([]modelApresentacao.ReqTasks, error)
}