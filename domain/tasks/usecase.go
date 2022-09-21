package tasks

import (
	"fmt"
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

func ListarStatusTasks(status string) (res []modelApresentacao.ReqTasks,err error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	dados, err := tasksRepo.ListarStatusTasks(status)
	
	if len(dados) == 0 {
		return res, fmt.Errorf("status does not exist")
	}
	return dados, err
}

func AtualizarTask(id string, req *modelApresentacao.ReqTask) (res *modelApresentacao.ReqTask, err error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)
	dados, err := tasksRepo.ListarTask(id)

	if err != nil {
		return res, fmt.Errorf("unable to update: Task not found in database")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}
	res, err = tasksRepo.AtualizarTask(id, req)
	if err != nil {
		return nil, fmt.Errorf("could not update: Team or Project does not exist")
	}
	return

}

func AtualizarStatusTask(id string, req *modelApresentacao.ReqTask) (res *modelApresentacao.ReqTask, err error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	res, err = tasksRepo.AtualizarStatusTask(id, req)
	if err != nil {
		return nil, err
	}	
	return
}

func DeletarTask(id string) (err error) {
	db := database.Conectar()
	defer db.Close()

	tasksRepo := tasks.NovoRepo(db)
	dados, err := tasksRepo.ListarTask(id)
	if err != nil {
		return fmt.Errorf("task does not exist")
	}

	if dados == nil {
		return fmt.Errorf("unrecognized error")
	}
	err = tasksRepo.DeletarTask(id)
	return
}