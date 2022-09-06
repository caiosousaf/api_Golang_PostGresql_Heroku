package projetos

import (
	"github.com/gin-gonic/gin"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
)

type IProjeto interface {
	NovoProjeto(req *modelApresentacao.ReqProjeto, c *gin.Context)
	ListarProjetos() ([]modelApresentacao.ReqProjetos, error)
	ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error)
	ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error)
}