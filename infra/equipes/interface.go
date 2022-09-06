package equipes

import (
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	"github.com/gin-gonic/gin"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context)
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
	BuscarMembrosDeEquipe(id string) ([]modelPessoa.ReqMembros, error)
	BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error)
	DeletarEquipe(id string) error
	AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error)
}