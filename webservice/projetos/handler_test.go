package projetos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/config/services"
	modelApresentacaoLogin "gerenciadorDeProjetos/domain/login/model"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
	"gerenciadorDeProjetos/webservice/login"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func GetToken() (token string){
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

func TestGetProjects(t *testing.T) {
	r := gin.Default()
	r.GET("/projetos/", ListarProjetos, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("Busca-Projetos-Sucesso", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/projetos/", nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		var projetos []modelApresentacao.ReqProjetos

		json.Unmarshal(w.Body.Bytes(), &projetos)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projetos)
	})
}
func TestGetProject(t *testing.T) {
	
	//Router(r.Group("/projetos", middlewares.Auth()))
	r := gin.Default()
	r.GET("/projetos/:id", ListarProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()

	t.Run("Busca-Projeto-Sucesso", func(t *testing.T) {
		
		id := "1"
		req, err := http.NewRequest("GET", "/projetos/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projetos)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projetos)

	})

	t.Run("Busca-Projeto-Erro-Id-Deve-Ser-Inteiro", func(t *testing.T) {
		
		id := "c"
		req, err := http.NewRequest("GET", "/projetos/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projetos)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, projetos)
	})

	t.Run("Busca-Projeto-Erro-Projeto-não-existe-com-o-id-passado", func(t *testing.T) {
		id := "6000"
		req, err := http.NewRequest("GET", "/projetos/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projetos)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, projetos)
	})

}

func TestGetTasksProject(t *testing.T) {
	r := gin.Default()
	r.GET("/projetos/:id/tasks", ListarTasksProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()

	t.Run("Busca-Task-Projeto-Sucesso", func(t *testing.T) {
		id := "1"
		req, err := http.NewRequest("GET", fmt.Sprintf("/projetos/%v/tasks", id), nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos []modelApresentacao.ReqTasksProjeto
		json.Unmarshal(w.Body.Bytes(),&projetos)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projetos)
	})

	t.Run("Busca-Task-Projeto-Error-ID-Invalido", func(t *testing.T) {
		id := "c"
		req, err := http.NewRequest("GET", fmt.Sprintf("/projetos/%v/tasks", id), nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos []modelApresentacao.ReqTasksProjeto
		json.Unmarshal(w.Body.Bytes(),&projetos)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, projetos)
	})

	t.Run("Busca-Task-Projeto-Erro-ID-Inexistente", func(t *testing.T) {
		id := "10000"
		req, err := http.NewRequest("GET", fmt.Sprintf("/projetos/%v/tasks", id), nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos []modelApresentacao.ReqTasksProjeto
		json.Unmarshal(w.Body.Bytes(),&projetos)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, projetos)
	})
}

func TestGetStatusOfAllProjects(t *testing.T) {
	r := gin.Default()
	r.GET("/projetos/status/:status", ListarProjetosComStatus, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()

	t.Run("Busca-Status-Projetos-Sucesso", func(t *testing.T) {
			status := "A Fazer"
			req, err := http.NewRequest("GET", "/projetos/status/"+status, nil)
			if err != nil {
				fmt.Println(err)
			}
	
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
			w := httptest.NewRecorder()
	
			r.ServeHTTP(w, req)
	
			var StatusProjects []modelApresentacao.ReqStatusProjeto
			json.Unmarshal(w.Body.Bytes(),&StatusProjects)
	
			assert.Equal(t, http.StatusOK, w.Code)
			assert.NotEmpty(t, StatusProjects)
	})

	t.Run("Busca-Status-Projetos-Erro-Status-Inexistente", func(t *testing.T) {
		status := "naoexiste"
		req, err := http.NewRequest("GET", "/projetos/status/"+status, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var StatusProjects []modelApresentacao.ReqStatusProjeto
		json.Unmarshal(w.Body.Bytes(),&StatusProjects)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, StatusProjects)
})
}