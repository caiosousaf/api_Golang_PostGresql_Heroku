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

	t.Run("Busca Projetos Sucesso", func(t *testing.T) {

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

	t.Run("BuscaProjetoSucesso", func(t *testing.T) {
		
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

	t.Run("BuscaProjetoErroIdDeveSerInteiro", func(t *testing.T) {
		
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
