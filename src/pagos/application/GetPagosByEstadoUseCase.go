// GetPagosByEstadoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagosByEstadoUseCase struct {
	repo repositories.IPagos
}

func NewGetPagosByEstadoUseCase(repo repositories.IPagos) *GetPagosByEstadoUseCase {
	return &GetPagosByEstadoUseCase{repo: repo}
}

func (uc *GetPagosByEstadoUseCase) Run(estado string) ([]entities.Pagos, error) {
	return uc.repo.GetByEstado(estado)
}