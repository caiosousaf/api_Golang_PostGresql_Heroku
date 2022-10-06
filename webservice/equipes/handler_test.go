package equipes

import (
	"bytes"
	"encoding/json"
	"fmt"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelApresentacaoLogin "gerenciadorDeProjetos/domain/login/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
	"gerenciadorDeProjetos/webservice/login"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/config/services"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
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

func TestGetTeams(t *testing.T) {
	r := gin.Default()
	r.GET("/equipes/", listarEquipes, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscarEquipesSucesso", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/equipes/", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipes []modelApresentacao.ReqEquipe
		json.Unmarshal(w.Body.Bytes(), &equipes)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.NotEmpty(t, equipes)
	})
}

func TestGetTeam(t *testing.T) {
	r := gin.Default()
	r.GET("/equipes/:id", buscarEquipe, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaEquipeSucesso", func(t *testing.T) {
		id := "1"
		req, _ := http.NewRequest("GET", "/equipes/"+id, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipe modelApresentacao.ReqEquipe
		json.Unmarshal(w.Body.Bytes(), &equipe)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, equipe)
	})

	t.Run("BuscaEquipeErroId", func(t *testing.T) {
		id := "5145"
		req, _ := http.NewRequest("GET", "/equipes/"+id, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipe modelApresentacao.ReqEquipe
		json.Unmarshal(w.Body.Bytes(), &equipe)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, equipe)
	})
}

func TestGetMembersTeam(t *testing.T) {
	r := gin.Default()
	r.GET("/equipes/:id/membros", buscarMembrosDeEquipe, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaMembrosSucesso", func(t *testing.T) {
		id := "1"
		req, _ := http.NewRequest("GET", fmt.Sprintf("/equipes/%v/membros", id), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipe []modelPessoa.ReqMembros
		json.Unmarshal(w.Body.Bytes(), &equipe)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, equipe)
	})

	t.Run("BuscaMembrosErroEquipe", func(t *testing.T) {
		id := "5145"
		req, _ := http.NewRequest("GET", fmt.Sprintf("/equipes/%v/membros", id), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipe []modelPessoa.ReqMembros
		json.Unmarshal(w.Body.Bytes(), &equipe)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, equipe)
	})

	t.Run("BuscaMembroSucessoSemMembros", func(t *testing.T) {
		id := "11"
		req, _ := http.NewRequest("GET", fmt.Sprintf("/equipes/%v/membros", id), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var equipe []modelPessoa.ReqMembros
		json.Unmarshal(w.Body.Bytes(), &equipe)

		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Empty(t, equipe)
	})
}
