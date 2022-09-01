package equipes

import (
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	"github.com/gin-gonic/gin"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context)
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
	BuscarMembrosDeEquipe(id string) ([]modelApresentacao.ReqEquipeMembros, error)
	BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error)
}