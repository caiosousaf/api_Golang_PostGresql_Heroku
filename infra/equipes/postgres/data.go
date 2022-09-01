package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelData "gerenciadorDeProjetos/infra/equipes/model"
	"github.com/gin-gonic/gin"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe, c *gin.Context) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Equipe)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("deu tudo certo")
}

func (postgres *DBEquipes) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe` // comando sql
	var res = []modelApresentacao.ReqEquipe{} // lista que vai receber resultados da consulta
	var equipe = modelApresentacao.ReqEquipe{} // estrutura individual que vai ser usada para preencher a lista

	rows, err := postgres.DB.Query(sqlStatement) // executando query e retornando as linhas e possíveis erros
	if err != nil {
		return nil, err // em caso de erro na consulta, a requisição será avortada e retornar um status 404 e o erro
	}
	for rows.Next() { // percorrendo linhas retornadas no sql
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil { // escaneando linha por linha e gravando na estrutura de equipe
			return nil, err // se houver algum erro, a função retorna ele
		}
		res = append(res, equipe) // preenchendo lista a cada iteração
	}
	fmt.Println("Listagem de todas as equipes deu certo!") // log que informa que essa parte geral deu certo
	return res, nil // retornando resposta do tipo []modelApresentacao.ReqEquipe
}

func (postgres *DBEquipes) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error){
	sqlStatement := `SELECT * FROM equipes WHERE id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	fmt.Println("Busca deu certo!")
	return equipe, nil
}

func (postgres *DBEquipes) BuscarMembrosDeEquipe(id string) ([]modelApresentacao.ReqEquipeMembros, error){
	sqlStatement := `select id_pessoa, nome_pessoa, funcao_pessoa, equipe_id, data_contratacao from pessoas WHERE equipe_id = $1`
	var res = []modelApresentacao.ReqEquipeMembros{}
	var equipe = modelApresentacao.ReqEquipeMembros{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.ID_Pessoa, &equipe.Nome_Pessoa, &equipe.Funcao_Pessoa, &equipe.Equipe_ID, &equipe.Data_Contratacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Busca de membros de uma equipe deu certo!")
	return res, nil
}

func (postgres *DBEquipes) BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error){
	sqlStatement := `select eq.nome_equipe, pr.id_projeto, pr.nome_projeto, pr.status, pr.descricao_projeto, pr.data_criacao, pr.data_conclusao, pr.prazo_entrega 
	from equipes as eq 
	inner join projetos as pr on eq.id_equipe = pr.equipe_id where eq.id_equipe = $1`
	var res = []modelApresentacao.ReqEquipeProjetos{}
	var equipe = modelApresentacao.ReqEquipeProjetos{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.Nome_Equipe, &equipe.ID_Projeto, &equipe.Nome_Projeto, &equipe.Status, &equipe.Descricao_Projeto,
			&equipe.Data_Criacao, &equipe.Data_Conclusao, &equipe.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Busca dos projetos de uma equipe deu certo!")
	return res, nil
}

func (postgres *DBEquipes) DeletarEquipe(id string) (*modelApresentacao.ReqEquipe, error){
	sqlStatement := `DELETE FROM equipes WHERE id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	rows := postgres.DB.QueryRow(sqlStatement, id)
	if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		return nil, nil
	}
	fmt.Println("Delete deu certo!")
	return equipe, nil
}

func (postgres *DBEquipes) AtualizarEquipe(id string, req *modelData.UpdateEquipe) (*modelApresentacao.ReqEquipe, error ){
	sqlStatement := `UPDATE equipes SET nome_equipe = $1 
	WHERE id_equipe = $2 RETURNING *`
	var equipe = &modelApresentacao.ReqEquipe{} 

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe, id)

	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		return nil, err
	}
	
	return equipe, nil
}