package pessoas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/config/services"
	modelApresentacaoLogin "gerenciadorDeProjetos/domain/login/model"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
	"gerenciadorDeProjetos/webservice/login"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

)

func TestGetPerson(t *testing.T) {

	r := gin.Default()

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
	Router(r.Group("pessoas", middlewares.Auth()) )

	t.Run("BuscaPessoasSucesso", func(t *testing.T) {
		
		req, _ := http.NewRequest("GET", "/pessoas", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoas modelApresentacao.ListarGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusOK, w.Code)
		
		assert.NotEmpty(t, pessoas)
	})

	t.Run("BuscaPessoaSucesso", func(t *testing.T) {
		//r.GET("/pessoas/:id", listarPessoa)
		id := "1"
		
		req, _ := http.NewRequest("GET", "/pessoas/"+id, nil)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoa modelApresentacao.ReqGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoa)

		assert.Equal(t, http.StatusOK, w.Code)
	
	})

}