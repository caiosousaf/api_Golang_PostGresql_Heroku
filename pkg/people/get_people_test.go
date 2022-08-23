package pessoas

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
func Test_handler_GetPeople(t *testing.T) {

	router := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(router, c)
	h := &handler{
		DB: c,
	}

	// Nome da rota ou nome do teste caso a rota precise de mais de um teste
	t.Run("BuscaPessoas", func(t *testing.T) {
		router.GET("/pessoas", h.GetPeople)
		req, _ := http.NewRequest("GET", "/pessoas", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var pessoas []GetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, pessoas)

	})

	t.Run("BuscaPessoa", func(t *testing.T) {

		w := httptest.NewRecorder()
		t.Run("BuscaPessoaSucesso", func(t *testing.T) {

			id := "4"
			router.GET("/pessoas/0", h.GetPerson)
			req, _ := http.NewRequest("GET", "/pessoas/"+id, nil)

			router.ServeHTTP(w, req)

			var pessoas GetPessoa
			json.Unmarshal(w.Body.Bytes(), &pessoas)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.NotEmpty(t, pessoas)
		})

		t.Run("ErrorBuscaPessoa", func(t *testing.T) {
			reqBadRequest, _ := http.NewRequest("GET", "/pessoas/600", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, reqBadRequest)
			assert.Equal(t, http.StatusBadRequest, w.Code)

		})

		t.Run("ErrorStatusNotFoundBuscaPessoa", func(t *testing.T) {

			reqNotFound, _ := http.NewRequest("GET", "/pessoas/c", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, reqNotFound)
			assert.Equal(t, http.StatusNotFound, w.Code)
		})
	})

	t.Run("BuscaTasksDePessoa", func(t *testing.T) {
		router.GET("/pessoas/0/tasks", h.GetPeople)
		req, _ := http.NewRequest("GET", "/pessoas/4/tasks", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var result []result

		json.Unmarshal(w.Body.Bytes(), &result)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, result)

		t.Run("BadRequestBuscaTasksDePessoa", func(t *testing.T) {
			reqBadRequest, _ := http.NewRequest("GET", "/pessoas/600/tasks", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, reqBadRequest)
			assert.Equal(t, http.StatusBadRequest, w.Code)

		})

		t.Run("StatusNotFoundBuscaTasksDePessoa", func(t *testing.T) {

			reqNotFound, _ := http.NewRequest("GET", "/pessoas/c/tasks", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, reqNotFound)
			assert.Equal(t, http.StatusNotFound, w.Code)
		})
	})

}

func Test_handler_PostPerson(t *testing.T) {
	router := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(router, c)
	h := &handler{
		DB: c,
	}

	router.POST("/pessoas", h.AddPerson)

	pessoa := AddPessoaRequestBody{
		Nome_Pessoa:   "Matheus Brisa",
		Funcao_Pessoa: "Back-End",
		Equipe_ID:     1,
	}
	jsonValue, _ := json.Marshal(pessoa)
	req, _ := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	t.Run("BadRequestBuscaTasksDePessoa", func(t *testing.T) {
		type AddPessoaWithError struct {
			Nome_Pessoa   string
			Funcao_Pessoa string
			Equipe_ID     string
		}
		pessoa := AddPessoaWithError{
			Nome_Pessoa:   "Matheus Brisa",
			Funcao_Pessoa: "",
			Equipe_ID:     "1",
		}
		jsonValue, _ := json.Marshal(pessoa)
		reqBadRequest, _ := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, reqBadRequest)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("StatusNotFoundBuscaTasksDePessoa", func(t *testing.T) {

		pessoa := AddPessoaRequestBody{
			Nome_Pessoa:   "Matheus Brisa",
			Funcao_Pessoa: "",
			Equipe_ID:     100,
		}
		jsonValue, _ := json.Marshal(pessoa)
		reqStatusNotFound, _ := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, reqStatusNotFound)
		assert.Equal(t, http.StatusNotFound, w.Code)

	})
}

func Test_handler_PutPerson(t *testing.T) {
	router := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(router, c)
	h := &handler{
		DB: c,
	}

	router.PUT("/pessoas/0", h.UpdatePerson)

	pessoa := UpdatePessoaRequestBody{
		Nome_Pessoa:   "Matheus Brisa",
		Funcao_Pessoa: "Back-End",
		Equipe_ID:     1,
	}
	jsonValue, _ := json.Marshal(pessoa)
	req, _ := http.NewRequest("PUT", "/pessoas/36", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	t.Run("BadRequestUpdatePessoa", func(t *testing.T) {
		type UpdatePessoaWithError struct {
			Nome_Pessoa   string
			Funcao_Pessoa string
			Equipe_ID     string
		}
		pessoa := UpdatePessoaWithError{
			Nome_Pessoa:   "Matheus Brisa",
			Funcao_Pessoa: "Front-End",
			Equipe_ID:     "1",
		}
		jsonValue, _ := json.Marshal(pessoa)
		reqBadRequest, _ := http.NewRequest("PUT", "/pessoas/36", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, reqBadRequest)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("NotFoundUpdatePessoa", func(t *testing.T) {

		pessoa := UpdatePessoaRequestBody{
			Nome_Pessoa:   "Matheus Brisa",
			Funcao_Pessoa: "Front-End",
			Equipe_ID:     1000,
		}
		jsonValue, _ := json.Marshal(pessoa)
		reqNotFound, _ := http.NewRequest("PUT", "/pessoas/36", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, reqNotFound)
		assert.Equal(t, http.StatusNotFound, w.Code)

	})

	t.Run("NotFoundUpdatePessoa", func(t *testing.T) {

		pessoa := UpdatePessoaRequestBody{
			Nome_Pessoa:   "Matheus Brisa",
			Funcao_Pessoa: "Front-End",
			Equipe_ID:     1,
		}
		jsonValue, _ := json.Marshal(pessoa)
		reqNotFound, _ := http.NewRequest("PUT", "/pessoas/3600", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, reqNotFound)
		assert.Equal(t, http.StatusNotFound, w.Code)

	})

}

func Test_handler_DeletePerson(t *testing.T) {
	router := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(router, c)
	h := &handler{
		DB: c,
	}

	router.DELETE("/pessoas/0", h.DeletePerson)

	w := httptest.NewRecorder()

	id := "86"

	req, _ := http.NewRequest("DELETE", "/pessoas/"+id, nil)

	router.ServeHTTP(w, req)

	var pessoas models.Pessoa
	json.Unmarshal(w.Body.Bytes(), &pessoas)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Run("BadRequestDeletePessoa", func(t *testing.T) {
		router.DELETE("/pessoas/00", h.DeletePerson)

		w := httptest.NewRecorder()

		id := "86"

		req, _ := http.NewRequest("DELETE", "/pessoas/"+id, nil)

		router.ServeHTTP(w, req)

		var pessoas models.Pessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}

func Test_handler_GetPersonName(t *testing.T) {
	router := gin.Default()

	dbUrl := "postgres://icsebrcphzbchf:02fde9fd34225b556aed45e81ca823f3c50b594f2530b3f95e8d2b1fe6517473@ec2-23-23-151-191.compute-1.amazonaws.com:5432/dcqvoffgfp6u50"
	c := db.Init(dbUrl)
	//n := r.Group("")
	//r.RouterGroup = *n

	RegisterRoutes(router, c)
	h := &handler{
		DB: c,
	}

	w := httptest.NewRecorder()

	id := "?person=caio%20swagger"
	router.GET("/pessoas/filtros/00", h.GetPersonName)
	req, _ := http.NewRequest("GET", "/pessoas/filtros/"+id, nil)

	router.ServeHTTP(w, req)

	var pessoas GetPessoa
	json.Unmarshal(w.Body.Bytes(), &pessoas)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Run("FiltroByFunction", func(t *testing.T) {
		id := "?person_function=Back-End"

		router.GET("/pessoas/filtros/01", h.GetPersonName)
		req, _ := http.NewRequest("GET", "/pessoas/filtros/"+id, nil)

		router.ServeHTTP(w, req)

		var pessoas []GetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
