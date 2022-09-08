package tasks

import (
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
)
type ITask interface {
	NovaTask(req *modelApresentacao.ReqTaskApresent) (*modelApresentacao.ReqTask, error)
	ListarTasks() ([]modelApresentacao.ReqTasks, error)
	ListarTask(id string) (*modelApresentacao.ReqTasks, error)
	ListarStatusTasks(status string) ([]modelApresentacao.ReqTasks, error)
}