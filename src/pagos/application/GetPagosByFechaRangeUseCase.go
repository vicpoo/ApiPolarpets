// GetPagosByFechaRangeUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagosByFechaRangeUseCase struct {
	repo repositories.IPagos
}

func NewGetPagosByFechaRangeUseCase(repo repositories.IPagos) *GetPagosByFechaRangeUseCase {
	return &GetPagosByFechaRangeUseCase{repo: repo}
}

func (uc *GetPagosByFechaRangeUseCase) Run(fechaInicio, fechaFin time.Time) ([]entities.Pagos, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}