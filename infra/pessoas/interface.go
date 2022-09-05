package pessoas

import (
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"github.com/gin-gonic/gin"
)
type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa, c *gin.Context)
	ListarPessoas() ([]modelApresentacao.ReqGetPessoa, error)
	ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error)
}