package pessoas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/config/services"
	modelApresentacaoLogin "gerenciadorDeProjetos/domain/login/model"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
	modelDataPessoa "gerenciadorDeProjetos/infra/pessoas/model"
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

func TestGetPeople(t *testing.T) {
	r := gin.Default()
	r.GET("/pessoas/", listarPessoas, middlewares.Auth())
	r.Use(cors.Default())
	//token := GetToken()

	t.Run("BuscaPessoasSucesso", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/pessoas/", nil)
		//req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoas modelApresentacao.ListarGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.NotEmpty(t, pessoas)
	})
}

func TestGetPerson(t *testing.T) {

	r := gin.Default()
	r.GET("/pessoas/:id", listarPessoa, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaPessoaSucesso", func(t *testing.T) {

		id := "1"
		req, err := http.NewRequest("GET", "/pessoas/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var pessoas modelApresentacao.ReqGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, pessoas)
	})

	t.Run("BuscaPessoaErroId", func(t *testing.T) {

		id := "2932"
		req, err := http.NewRequest("GET", "/pessoas/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var pessoas modelApresentacao.ReqGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoas)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, pessoas)
	})
}

func TestGetTasksPerson(t *testing.T) {
	r := gin.Default()
	r.GET("/pessoas/:id/tasks", listarTarefasPessoa, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("BuscaTaskPessoaSucesso", func(t *testing.T) {

		id := "1"
		req, err := http.NewRequest("GET", fmt.Sprintf("/pessoas/%v/tasks", id), nil)

		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var tasks []modelApresentacao.ReqTarefaPessoa
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, tasks)
	})

	t.Run("BuscaTaskPessoaErroSemTasks", func(t *testing.T) {

		id := "11"
		req, err := http.NewRequest("GET", fmt.Sprintf("/pessoas/%v/tasks", id), nil)

		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var tasks []modelApresentacao.ReqTarefaPessoa
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Empty(t, tasks)
	})

	t.Run("BuscaTaskPessoaInexistente", func(t *testing.T) {

		id := "1145"
		req, err := http.NewRequest("GET", fmt.Sprintf("/pessoas/%v/tasks", id), nil)

		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var tasks []modelApresentacao.ReqTarefaPessoa
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, tasks)
	})
}

func TestAddPerson(t *testing.T) {
	r := gin.Default()
	r.POST("/pessoas", novaPessoa, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("AdicionarPessoaSucesso", func(t *testing.T) {

		nome_pessoa := "Teste Unitarios"
		funcao_pessoa := "Full-Stack"
		equipe := 10

		pessoa := modelDataPessoa.ReqPessoa{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		req, err := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAdicionada modelApresentacao.ReqPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAdicionada)
		

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.NotEmpty(t, pessoaAdicionada)
	})

	t.Run("AdicionarPessoaErroParametros", func(t *testing.T) {

		nome_pessoa := "Teste Unitario"
		funcao_pessoa := "Full-Stack"
		equipe := "10"

		type reqPessoaForcaError struct {
			Nome_Pessoa   *string `json:"nome_pessoa"`
			Funcao_Pessoa *string `json:"funcao_pessoa"`
			Equipe_ID     *string    `json:"equipe_id" `
		}

		pessoa := reqPessoaForcaError{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		req, err := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAdicionada modelApresentacao.ReqPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAdicionada)
		

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.Empty(t, pessoaAdicionada)
	})

	t.Run("AdicionarPessoaErroEquipeInexistente", func(t *testing.T) {

		nome_pessoa := "Teste Unitario"
		funcao_pessoa := "Full-Stack"
		equipe := 1052

		pessoa := modelDataPessoa.ReqPessoa{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		req, err := http.NewRequest("POST", "/pessoas", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAdicionada modelApresentacao.ReqPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAdicionada)
		

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.Empty(t, pessoaAdicionada)
	})
}

func TestUpdatePerson(t *testing.T) {
	r := gin.Default()
	r.PUT("/pessoas/:id", atualizarPessoa, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("AtualizarSucesso", func(t *testing.T) {
		nome_pessoa := "OI"
		funcao_pessoa := "Opa"
		equipe := 10

		pessoa := modelDataPessoa.ReqPessoa{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		id := "11"
		req, err := http.NewRequest("PUT", "/pessoas/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAtualizada modelApresentacao.ReqAtualizarPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAtualizada)
		
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.NotEmpty(t, pessoaAtualizada)
	})

	t.Run("AtualizarErroParametros", func(t *testing.T) {
		nome_pessoa := "OI"
		funcao_pessoa := "Opa"
		equipe := "10"

		type reqPessoaForcaError struct {
			Nome_Pessoa   *string `json:"nome_pessoa"`
			Funcao_Pessoa *string `json:"funcao_pessoa"`
			Equipe_ID     *string    `json:"equipe_id" `
		}

		pessoa := reqPessoaForcaError{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		id := "11"
		req, err := http.NewRequest("PUT", "/pessoas/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAtualizada modelApresentacao.ReqAtualizarPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAtualizada)
		
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.Empty(t, pessoaAtualizada)
	})

	t.Run("AtualizarErroIdInexistente", func(t *testing.T) {
		nome_pessoa := "OI"
		funcao_pessoa := "Opa"
		equipe := 10

		pessoa := modelDataPessoa.ReqPessoa{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		id := "111451"
		req, err := http.NewRequest("PUT", "/pessoas/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAtualizada modelApresentacao.ReqAtualizarPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAtualizada)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.Empty(t, pessoaAtualizada)
	})

	t.Run("AtualizarEquipeInexistente", func(t *testing.T) {
		nome_pessoa := "OI"
		funcao_pessoa := "Opa"
		equipe := 1045445

		pessoa := modelDataPessoa.ReqPessoa{
			Nome_Pessoa:   &nome_pessoa,
			Funcao_Pessoa: &funcao_pessoa,
			Equipe_ID:     &equipe,
		}

		jsonValue, _ := json.Marshal(pessoa)
		id := "11"
		req, err := http.NewRequest("PUT", "/pessoas/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var pessoaAtualizada modelApresentacao.ReqAtualizarPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoaAtualizada)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NotEmpty(t, pessoa)
		assert.Empty(t, pessoaAtualizada)
	})
}

func TestDeletePerson(t *testing.T) {
	r := gin.Default()
	r.DELETE("/pessoas/:id", deletarPessoa, middlewares.Auth())
	r.Use(cors.Default())
	token := GetToken()

	t.Run("DeletarSucesso", func(t *testing.T) {
		id := "26"
		req, err := http.NewRequest("DELETE", "/pessoas/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("DeletarErroIdInexistente", func(t *testing.T) {
		id := "12151"
		req, err := http.NewRequest("DELETE", "/pessoas/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetFilterPerson(t *testing.T) {
	r := gin.Default()
	r.GET("/pessoas/filtros", listarPessoasFiltro, middlewares.Auth())
	
	r.Use(cors.Default())
	token := GetToken()

	t.Run("FiltroPessoaSucesso", func(t *testing.T) {
		
		req, err := http.NewRequest("GET", "/pessoas/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("order", "ca")
		q.Add("orderBy", "nome_pessoa")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
   
		var pessoa modelApresentacao.ListarGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoa)
 
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, pessoa)
	})	

	t.Run("FiltroPessoaSucesso", func(t *testing.T) {
		
		req, err := http.NewRequest("GET", "/pessoas/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
   
		var pessoa modelApresentacao.ListarGetPessoa
		json.Unmarshal(w.Body.Bytes(), &pessoa)
 
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, pessoa)
	})	
}