package projetos

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/db"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

//go test -v -cover ./...
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
func Test_handler_GetProjects(t *testing.T) {
	r := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(r, c)
	h := &handler{
		DB: c,
	}

	t.Run("BuscaProjetosSucesso", func(t *testing.T) {
		r.GET("/projetos", h.GetProjects)
		req, _ := http.NewRequest("GET", "/projetos", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetos []Projects
		json.Unmarshal(w.Body.Bytes(), &projetos)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projetos)

	})

}

func Test_handler_GetProject(t *testing.T) {
	r := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(r, c)
	h := &handler{
		DB: c,
	}

	r.GET("/projetos/0", h.GetProject)

	t.Run("SucessoBuscaProjeto", func(t *testing.T) {
		id := "92"
		req, _ := http.NewRequest("GET", "/projetos/"+id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projeto []Projects
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)

	})

	t.Run("StatusNotFoundBuscaProjeto", func(t *testing.T) {

		id := "c"
		req, _ := http.NewRequest("GET", "/projetos/"+id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projeto Projects
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusNotFound, w.Code)

	})

	t.Run("StatusBadRequestBuscaProjeto", func(t *testing.T) {

		id := "10000"
		req, _ := http.NewRequest("GET", "/projetos/"+id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projeto []Projects
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}

func Test_handler_GetTasksProject(t *testing.T) {
	r := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(r, c)
	h := &handler{
		DB: c,
	}

	r.GET("/projetos/0/tasks", h.GetProjectTasks)

	t.Run("SucessoBuscaTasksDeProjeto", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/projetos/92/tasks", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var tasks []TasksProjeto
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, tasks)

	})

	t.Run("BadRequestBuscaTasksDeProjeto", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/projetos/9200/tasks", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var tasks []TasksProjeto
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("StatusNotFoundBuscaTasksDeProjeto", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/projetos/c/tasks", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var tasks []TasksProjeto
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusNotFound, w.Code)

	})
}

func Test_handler_AddProject(t *testing.T) {
	r := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(r, c)
	h := &handler{
		DB: c,
	}

	r.POST("/projetos", h.AddProject)

	// Esse post funciona pois não existe nenhum projeto cadastrado com esse nome
	t.Run("SucessoPostDeProjeto", func(t *testing.T) {

		projeto := AddProjetoRequestBody{
			Nome_Projeto:      "Teste Unitário",
			Descricao_Projeto: "Descrição de Teste Unitário",
			Equipe_ID:         2,
			Prazo:             10,
		}

		jsonValue, _ := json.Marshal(projeto)
		req, _ := http.NewRequest("POST", "/projetos", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	// Não é possivel criar um projeto com o mesmo nome de outro projeto que já esteja cadastrado
	t.Run("StatusNotFoundPostDeProjeto", func(t *testing.T) {

		projeto := AddProjetoRequestBody{
			Nome_Projeto:      "Teste Unitário",
			Descricao_Projeto: "Descrição de Teste Unitário",
			Equipe_ID:         2,
			Prazo:             10,
		}

		jsonValue, _ := json.Marshal(projeto)
		req, _ := http.NewRequest("POST", "/projetos", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

}

func Test_handler_DeleteProject(t *testing.T) {
	r := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(r, c)
	h := &handler{
		DB: c,
	}

	r.DELETE("/projetos", h.DeleteProject)

	t.Run("SucessoDeleteDeProjeto", func(t *testing.T) {
		w := httptest.NewRecorder()

		id := "121"

		req, _ := http.NewRequest("DELETE", "/projetos/"+id, nil)

		r.ServeHTTP(w, req)

		var projeto models.Projeto
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("BadRequestDeleteDeProjeto", func(t *testing.T) {
		w := httptest.NewRecorder()

		id := "121"

		req, _ := http.NewRequest("DELETE", "/projetos/"+id, nil)

		r.ServeHTTP(w, req)

		var projeto models.Projeto
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("StatusNotFoundDeleteDeProjeto", func(t *testing.T) {
		w := httptest.NewRecorder()

		id := "c"

		req, _ := http.NewRequest("DELETE", "/projetos/"+id, nil)

		r.ServeHTTP(w, req)

		var projeto models.Projeto
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
