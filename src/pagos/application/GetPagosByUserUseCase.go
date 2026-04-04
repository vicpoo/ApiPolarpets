// GetPagosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagosByUserUseCase struct {
	repo repositories.IPagos
}

func NewGetPagosByUserUseCase(repo repositories.IPagos) *GetPagosByUserUseCase {
	return &GetPagosByUserUseCase{repo: repo}
}

func (uc *GetPagosByUserUseCase) Run(idUsuario int32) ([]entities.Pagos, error) {
	return uc.repo.GetByUser(idUsuario)
}