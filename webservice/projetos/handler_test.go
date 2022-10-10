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
	modelDataProjetos "gerenciadorDeProjetos/infra/projetos/model"
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

func GetId() (id uint) {
	var t = &testing.T{}
	r := gin.Default()
	r.Use(cors.Default())
	token := GetToken()

	r.GET("/projetos/filtros", listarProjetosFiltro)
	req, err := http.NewRequest("GET", "/projetos/filtros", nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("order", "desc")
	q.Add("orderBy", "id_projeto")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var projeto []modelApresentacao.ReqProjetos
	json.Unmarshal(w.Body.Bytes(), &projeto)
	id = *projeto[0].ID_Projeto

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, projeto)

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

	id := fmt.Sprint(GetId())
	
	t.Run("Busca-Projeto-Sucesso", func(t *testing.T) {

		
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
	r.Use(middlewares.Auth())

	token := GetToken()
	id := fmt.Sprint(GetId())

	t.Run("Busca-Task-Projeto-Sucesso", func(t *testing.T) {
		
		req, err := http.NewRequest("GET", fmt.Sprintf("/projetos/%v/tasks", id), nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projetos []modelApresentacao.ReqTasksProjeto
		json.Unmarshal(w.Body.Bytes(), &projetos)

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
		json.Unmarshal(w.Body.Bytes(), &projetos)

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
		json.Unmarshal(w.Body.Bytes(), &projetos)

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
		json.Unmarshal(w.Body.Bytes(), &StatusProjects)

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
		json.Unmarshal(w.Body.Bytes(), &StatusProjects)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, StatusProjects)
	})
}

func TestAddProject(t *testing.T) {
	r := gin.Default()
	r.POST("/projetos", NovoProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()

	t.Run("post-sucesso", func(t *testing.T) {

		nome_projeto := "Teste sjdnsajkcnsssadasasda saffsdfddd dsdsadscsac xxssxa dvvbsdjk"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe := 1

		projeto := modelDataProjetos.ReqProjeto{
			Nome_Projeto:      &nome_projeto,
			Descricao_Projeto: &descricao_projeto,
			Equipe_ID:         &equipe,
			Prazo:             2,
		}

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("POST", "/projetos", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAdicionado modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projetoAdicionado)
		v := projetoAdicionado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.NotEmpty(t, projeto)
		assert.NotEmpty(t, projetoAdicionado)
	})

	t.Run("post-fail-projeto-ja-existe", func(t *testing.T) {
		nome_projeto := "Teste sjdnsajkcnsssadasasda saffsdfddd dsdsadscsac xxssxa dvvbsdjk"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe := 1

		projeto := modelDataProjetos.ReqProjeto{
			Nome_Projeto:      &nome_projeto,
			Descricao_Projeto: &descricao_projeto,
			Equipe_ID:         &equipe,
			Prazo:             2,
		}

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("POST", "/projetos", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("post-fail-parametros-errados", func(t *testing.T) {
		type ReqProjetoForcaError struct {
			Nome_Projeto      *string `json:"nome_projeto"`
			Equipe_ID         *string `json:"equipe_id"`
			Descricao_Projeto *string `json:"descricao_projeto"`
			Prazo             int     `json:"prazo_entrega"`
		}
		nome_projeto := "Teste sjdnsajkcnsssadasasda saffsdfddd dsdsadscsac xxssxa dvvbsdjk"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe := "1"

		projeto := ReqProjetoForcaError{
			Nome_Projeto:      &nome_projeto,
			Descricao_Projeto: &descricao_projeto,
			Equipe_ID:         &equipe,
			Prazo:             2,
		}

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("POST", "/projetos", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}

func TestUpdateProject(t *testing.T) {
	r := gin.Default()
	r.PUT("/projetos/:id", AtualizarProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())
	t.Run("PUT-sucesso", func(t *testing.T) {
		nome_projeto := "Atualiza ai"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe_id := 1

		projeto := modelDataProjetos.ReqAtualizarProjetoData{
			Nome_Projeto:      &nome_projeto,
			Equipe_ID:         &equipe_id,
			Descricao_Projeto: &descricao_projeto,
		}

		jsonValue, _ := json.Marshal(projeto)
		
		req, err := http.NewRequest("PUT", "/projetos/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)
		assert.NotEmpty(t, projetoAtualizado)
	})

	t.Run("PUT-erro-parametros", func(t *testing.T) {
		nome_projeto := "Teste sjdnsajkcnsssadasasda saffsdfddd dsdsadscsac xxssxa dvvbsdjk"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe := "1"

		type reqAtualizarProjetoDataForcaError struct {
			Nome_Projeto      *string `json:"nome_projeto" example:"Casas Bahias"`
			Equipe_ID         *string `json:"equipe_id" example:"1"`
			Descricao_Projeto *string `json:"descricao_projeto" example:"Criacao de sistema e-commerce"`
		}

		projeto := reqAtualizarProjetoDataForcaError{
			Nome_Projeto:      &nome_projeto,
			Equipe_ID:         &equipe,
			Descricao_Projeto: &descricao_projeto,
		}

		jsonValue, _ := json.Marshal(projeto)
		id := fmt.Sprint(GetId())
		req, err := http.NewRequest("PUT", "/projetos/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)
		v := projetoAtualizado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, projeto)
		assert.Empty(t, projetoAtualizado)
	})

	t.Run("PUT-erro-ID", func(t *testing.T) {
		nome_projeto := "Teste sjdnsajkcnsssadasasda saffsdfddd dsdsadscsac xxssxa dvvbsdjk"
		descricao_projeto := "Descricao teste para o teste unitário testarr"
		equipe := 1

		projeto := modelDataProjetos.ReqAtualizarProjetoData{
			Nome_Projeto:      &nome_projeto,
			Equipe_ID:         &equipe,
			Descricao_Projeto: &descricao_projeto,
		}

		jsonValue, _ := json.Marshal(projeto)
		id := "104141"
		req, err := http.NewRequest("PUT", "/projetos/"+id, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)
		v := projetoAtualizado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NotEmpty(t, projeto)
		assert.Empty(t, projetoAtualizado)
	})
}

func TestUpdateStatus(t *testing.T) {
	r := gin.Default()
	r.PUT("/projetos/:id/status", AtualizarStatusProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())
	t.Run("PUT-sucesso", func(t *testing.T) {
		status := "Em Andamento"

		projeto := modelDataProjetos.ReqUpdateStatusProjeto{
			Status: &status,
		}
		

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/projetos/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)
		v := projetoAtualizado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)
		assert.NotEmpty(t, projetoAtualizado)
	})

	t.Run("PUT-erro-parametros", func(t *testing.T) {
		status := 1

		type ReqUpdateStatusProjetoForcaError struct {
			Status *int `json:"status" example:"Em Andamento"`
		}

		projeto := ReqUpdateStatusProjetoForcaError{
			Status: &status,
		}
		id := fmt.Sprint(GetId())

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/projetos/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)
		v := projetoAtualizado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, projeto)
		assert.Empty(t, projetoAtualizado)
	})

	t.Run("PUT-sucesso", func(t *testing.T) {
		status := "Em Andamento"

		projeto := modelDataProjetos.ReqUpdateStatusProjeto{
			Status: &status,
		}
		id := "10151"

		jsonValue, _ := json.Marshal(projeto)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/projetos/%v/status", id), bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var projetoAtualizado modelApresentacao.ReqAtualizarProjeto
		json.Unmarshal(w.Body.Bytes(), &projetoAtualizado)
		v := projetoAtualizado.ID_Projeto
		fmt.Println(v)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NotEmpty(t, projeto)
		assert.Empty(t, projetoAtualizado)
	})
}

func TestDeleteProject(t *testing.T) {
	r := gin.Default()
	r.DELETE("/projetos/:id", DeletarProjeto, middlewares.Auth())
	r.Use(cors.Default())

	token := GetToken()
	id := fmt.Sprint(GetId())

	t.Run("delete-project-sucesso", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/projetos/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("delete-project-erro", func(t *testing.T) {
		id := "12151"
		req, err := http.NewRequest("DELETE", "/projetos/"+id, nil)
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
	r.GET("/projetos/filtros", listarProjetosFiltro, middlewares.Auth())

	r.Use(cors.Default())
	token := GetToken()

	t.Run("FiltroProjetoSucesso", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/projetos/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("value", "bl")
		q.Add("column", "nome_projeto")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projeto []modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)
	})

	t.Run("FiltroProjetoSucessoOrder", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/projetos/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}
		q := req.URL.Query()
		q.Add("order", "desc")
		q.Add("orderBy", "id_projeto")
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projeto []modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projeto)
		opa := *projeto[0].ID_Projeto
		fmt.Println(opa)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)
	})

	t.Run("FiltroProjetoSucessoSemQuery", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/projetos/filtros", nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		var projeto []modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projeto)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, projeto)
	})

	t.Run("FiltroPessoaSucesso", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/projetos/filtros", nil)
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

		var projeto modelApresentacao.ReqProjetos
		json.Unmarshal(w.Body.Bytes(), &projeto)
		// opa := *pessoa.Pessoas[0].ID_Pessoa
		// fmt.Println(opa)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, projeto)
	})
}
