package tasks

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	"gerenciadorDeProjetos/infra/tasks"
)

func NovaTask(req *modelApresentacao.ReqTaskApresent) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.NovaTask(req)
}

func ListarTasks() ([]modelApresentacao.ReqTasks, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.ListarTasks()
}

func ListarTask(id string) (*modelApresentacao.ReqTasks, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.ListarTask(id)
}

func ListarStatusTasks(status string) ([]modelApresentacao.ReqTasks, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.ListarStatusTasks(status)
}

func AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.AtualizarTask(id, req)
}