// GetRegistroHabitoByFechaRangeUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetRegistroHabitoByFechaRangeUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroHabitoByFechaRangeUseCase(repo repositories.IRegistroHabito) *GetRegistroHabitoByFechaRangeUseCase {
	return &GetRegistroHabitoByFechaRangeUseCase{repo: repo}
}

func (uc *GetRegistroHabitoByFechaRangeUseCase) Run(fechaInicio, fechaFin time.Time) ([]entities.RegistroHabito, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}