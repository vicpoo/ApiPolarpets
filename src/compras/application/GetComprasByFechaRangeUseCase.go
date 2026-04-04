// GetComprasByFechaRangeUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetComprasByFechaRangeUseCase struct {
	repo repositories.ICompras
}

func NewGetComprasByFechaRangeUseCase(repo repositories.ICompras) *GetComprasByFechaRangeUseCase {
	return &GetComprasByFechaRangeUseCase{repo: repo}
}

func (uc *GetComprasByFechaRangeUseCase) Run(fechaInicio, fechaFin time.Time) ([]entities.Compras, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}