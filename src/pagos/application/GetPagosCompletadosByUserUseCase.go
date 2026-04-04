// GetPagosCompletadosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagosCompletadosByUserUseCase struct {
	repo repositories.IPagos
}

func NewGetPagosCompletadosByUserUseCase(repo repositories.IPagos) *GetPagosCompletadosByUserUseCase {
	return &GetPagosCompletadosByUserUseCase{repo: repo}
}

func (uc *GetPagosCompletadosByUserUseCase) Run(idUsuario int32) ([]entities.Pagos, error) {
	return uc.repo.GetPagosCompletadosByUser(idUsuario)
}