package tasks

import (
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
)
type ITask interface {
	NovaTask(req *modelApresentacao.ReqTaskApresent) (*modelApresentacao.ReqTask, error)
}