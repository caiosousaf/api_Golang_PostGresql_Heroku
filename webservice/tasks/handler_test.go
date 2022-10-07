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
	//modelDataTasks "gerenciadorDeProjetos/infra/tasks/model"
	"gerenciadorDeProjetos/webservice/login"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	t.Run("BuscaTaskSucesso", func(t *testing.T) {
		id := "1"
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



