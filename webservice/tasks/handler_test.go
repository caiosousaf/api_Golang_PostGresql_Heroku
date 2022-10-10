package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/config/services"
	modelApresentacaoLogin "gerenciadorDeProjetos/domain/login/model"
	modelApresentacao "gerenciadorDeProjetos/domain/tasks/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
	modelDataTasks "gerenciadorDeProjetos/infra/tasks/model"
	"gerenciadorDeProjetos/webservice/login"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func GetToken() (token string) {
	var t = &testing.T{}
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/login", login.Login)
	usuario := modelData.ReqLogin{
		Email:    "caio@caio.com",
		Password: "salmo",
	}
	jsonValue, _ := json.Marshal(usuario)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	pessoa := modelApresentacaoLogin.Login{}

	token, err := services.NewJWTService().GenerateToken(pessoa.ID_Usuario)
	if err != nil {
		return
	}
	fmt.Println(token)
	return
}

func GetId() (id uint) {
	var t = &testing.T{}
	r := gin.Default()
	r.Use(cors.Default())
	token := GetToken()

	r.GET("/tasks/filtros", listarTasksFiltro)
	req, err := http.NewRequest("GET", "/tasks/filtros", nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("order", "desc")
	q.Add("orderBy", "id_task")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var task []modelApresentacao.ReqTasks
	json.Unmarshal(w.Body.Bytes(), &task)
	id = *task[0].ID_Task
	

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, task)

	return
}

func TestAddTask(t *testing.T) {
	r := gin.Default()
	r.POST("/tasks", NovaTask, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("POST-sucesso", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoa_id := 10
		projeto_id := 1
		prazo := 2
		prioridade := 1

		task := modelDataTasks.ReqTaskData {
			Descricao_Task: &descricao_task,
			PessoaID: &pessoa_id,
			ProjetoID: &projeto_id,
			Prazo: prazo,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAdicionada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAdicionada)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.NotEmpty(t, task)
		assert.NotEmpty(t, taskAdicionada)
	})

	t.Run("POST-Erro-Parametros", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoa_id := 10
		projeto_id := "1"
		prazo := 2
		prioridade := 1

		type reqTaskDataForcaError struct {
			Descricao_Task *string `json:"descricao_task" example:"Descrição Teste"`
			PessoaID       *int    `json:"pessoa_id" example:"4"`
			ProjetoID      *string    `json:"projeto_id" example:"24"`
			Prazo          int     `json:"prazo_entrega" example:"17"`
			Prioridade     *int    `json:"prioridade" example:"1"`
		}

		task := reqTaskDataForcaError {
			Descricao_Task: &descricao_task,
			PessoaID: &pessoa_id,
			ProjetoID: &projeto_id,
			Prazo: prazo,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAdicionada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAdicionada)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAdicionada)
	})

	t.Run("POST-erro-Id-pessoa", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoa_id := 10541
		projeto_id := 1
		prazo := 2
		prioridade := 1

		task := modelDataTasks.ReqTaskData {
			Descricao_Task: &descricao_task,
			PessoaID: &pessoa_id,
			ProjetoID: &projeto_id,
			Prazo: prazo,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAdicionada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAdicionada)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAdicionada)
	})
}

func TestGetTasks(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks/", ListarTasks, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaTasksSucesso", func(t *testing.T) {
	
			req, err := http.NewRequest("GET", "/tasks/", nil)
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
			w := httptest.NewRecorder()
	
			r.ServeHTTP(w, req)
			var tasks []modelApresentacao.ReqTasks
	
			json.Unmarshal(w.Body.Bytes(), &tasks)
	
			assert.Equal(t, http.StatusOK, w.Code)
			assert.NotEmpty(t, tasks)
	})
}

func TestGetTask(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks/:id", ListarTask, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()
	id := fmt.Sprint(GetId())
	t.Run("BuscaTaskSucesso", func(t *testing.T) {
		
		req, err := http.NewRequest("GET", "/tasks/"+id, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		var task modelApresentacao.ReqTask

		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)
	})

	t.Run("BuscaTaskErroId", func(t *testing.T) {
		id := "154154"
		req, err := http.NewRequest("GET", "/tasks/"+id, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		var task modelApresentacao.ReqTask

		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, task)
	})
}

func TestGetStatusTask(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks/status/:status", ListarStatusTasks, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaStatusSucesso", func(t *testing.T) {
			status := "A Fazer"
			req, err := http.NewRequest("GET", "/tasks/status/"+status, nil)
	
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
			w := httptest.NewRecorder()
	
			r.ServeHTTP(w, req)
			var tasks []modelApresentacao.ReqTasks
	
			json.Unmarshal(w.Body.Bytes(), &tasks)
	
			assert.Equal(t, http.StatusOK, w.Code)
			assert.NotEmpty(t, tasks)
	})

	t.Run("BuscaStatusErroStatus", func(t *testing.T) {
		status := 1
		req, err := http.NewRequest("GET", fmt.Sprintf("/tasks/status/%v", status), nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		var tasks []modelApresentacao.ReqTasks

		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, tasks)
})
}

func TestUpdateTask(t *testing.T) {
	r := gin.Default()
	r.PUT("/tasks/:id", AtualizarTask, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())

	t.Run("PUT-sucesso", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoa_id := 10
		projeto_id := 1
		prioridade := 1

		task := modelDataTasks.ReqUpdateTaskData{
			Descricao_Task: &descricao_task,
			PessoaID: &pessoa_id,
			ProjetoID: &projeto_id,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)

		req, err := http.NewRequest("PUT", "/tasks/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var tasksAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &tasksAtualizada)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)
		assert.NotEmpty(t, tasksAtualizada)
	})

	t.Run("PUT-erro-id", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoaid := 10
		projetoid := 1
		prioridade := 1

		task := modelDataTasks.ReqTaskData{
			Descricao_Task: &descricao_task,
			PessoaID: &pessoaid,
			ProjetoID: &projetoid,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)
		idErro := "10155"
		req, err := http.NewRequest("PUT", "/tasks/"+idErro, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAtualizada)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAtualizada)
	})

	t.Run("PUT-erro-parametro", func(t *testing.T) {
		descricao_task := uuid.New().String()
		pessoaid := "10"
		projetoid := 1
		prioridade := 1

		type ReqUpdateTaskDataForcaError struct {
			Descricao_Task *string `json:"descricao_task" example:"Descrição Teste"`
			PessoaID       *string    `json:"pessoa_id" example:"4"`
			ProjetoID      *int    `json:"projeto_id" example:"24"`
			Prioridade     *int    `json:"prioridade" example:"1"`
		}

		task := ReqUpdateTaskDataForcaError{
			Descricao_Task: &descricao_task,
			PessoaID: &pessoaid,
			ProjetoID: &projetoid,
			Prioridade: &prioridade,
		}

		jsonValue, _ := json.Marshal(task)
		id := "10155"
		req, err := http.NewRequest("PUT", "/tasks/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAtualizada)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAtualizada)
	})
}

func TestUpdateStatusTask(t *testing.T) {
	r := gin.Default()
	r.PUT("/tasks/:id/status", AtualizarStatusTask, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())

	t.Run("AtualizarStatusSucesso", func(t *testing.T) {
		status := "Em Andamento"

		task := modelDataTasks.ReqUpdateStatusTask{
			Status: &status,
		}
		

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/tasks/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAtualizada)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)
		assert.NotEmpty(t, taskAtualizada)
	})

	t.Run("AtualizarStatusErroId", func(t *testing.T) {
		status := "Em Andamento"

		task := modelDataTasks.ReqUpdateStatusTask{
			Status: &status,
		}
		id := "101414"

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/tasks/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAtualizada)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAtualizada)
	})

	t.Run("AtualizarStatusErroStatus", func(t *testing.T) {
		status := 1

		type reqUpdateStatusTaskForcaErro struct {
			Status *int `json:"status" example:"Em Teste"`
		}

		task := reqUpdateStatusTaskForcaErro{
			Status: &status,
		}
		

		jsonValue, _ := json.Marshal(task)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/tasks/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var taskAtualizada modelApresentacao.ReqTask
		json.Unmarshal(w.Body.Bytes(), &taskAtualizada)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, task)
		assert.Empty(t, taskAtualizada)
	})
}

func TestDeleteTask(t *testing.T) {
	r := gin.Default()
	r.DELETE("/tasks/:id", DeletarTask, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())

	t.Run("delete-task-sucesso", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("delete-task-erro-id", func(t *testing.T) {
		id := "12151"
		req, err := http.NewRequest("DELETE", "/tasks/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

	})
}

func TestGetFilterProject(t *testing.T) {
	r := gin.Default()
	r.GET("/tasks/filtros", listarTasksFiltro, middlewares.Auth())

	r.Use(cors.Default())
	token := GetToken()

	t.Run("FiltroTaskSucesso", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/tasks/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("value", "p")
		q.Add("column", "descricao_task")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var task []modelApresentacao.ReqTasks
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)
	})

	t.Run("FiltroTaskSucessoOrder", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/tasks/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("order", "desc")
		q.Add("orderBy", "id_task")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var task []modelApresentacao.ReqTasks
		json.Unmarshal(w.Body.Bytes(), &task)
		opa := *task[0].ID_Task
		fmt.Println(opa)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)
	})

	t.Run("FiltroTaskSucessoSemQuery", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/tasks/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var tasks []modelApresentacao.ReqTasks
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, tasks)
	})

	t.Run("FiltroTaskSucesso", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/tasks/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("order", "asl")
		q.Add("orderBy", "equipeid")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var task modelApresentacao.ReqTasks
		json.Unmarshal(w.Body.Bytes(), &task)
		// opa := *pessoa.Pessoas[0].ID_Pessoa
		// fmt.Println(opa)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, task)
	})
}