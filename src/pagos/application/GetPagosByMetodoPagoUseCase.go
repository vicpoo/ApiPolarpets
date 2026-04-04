// GetPagosByMetodoPagoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagosByMetodoPagoUseCase struct {
	repo repositories.IPagos
}

func NewGetPagosByMetodoPagoUseCase(repo repositories.IPagos) *GetPagosByMetodoPagoUseCase {
	return &GetPagosByMetodoPagoUseCase{repo: repo}
}

func (uc *GetPagosByMetodoPagoUseCase) Run(metodoPago string) ([]entities.Pagos, error) {
	return uc.repo.GetByMetodoPago(metodoPago)
}