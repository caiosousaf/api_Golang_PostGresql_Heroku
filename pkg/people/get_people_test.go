package pessoas

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/db"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

//go test -v -cover ./...
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html 
func  Test_handler_GetPeople(t *testing.T) {


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

			id  := "4"
			router.GET("/pessoas/0", h.GetPerson)
			req, _ := http.NewRequest("GET", "/pessoas/" + id  , nil)

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
		
		
		
		t.Run("ErrorBadGatewayBuscaPessoa", func(t *testing.T) { 
			
			reqNotFound, _ := http.NewRequest("GET", "/pessoas/c" , nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, reqNotFound)
			assert.Equal(t, http.StatusNotFound, w.Code)
		})
	})
	
}
