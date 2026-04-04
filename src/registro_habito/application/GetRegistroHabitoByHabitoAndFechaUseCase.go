// GetRegistroHabitoByHabitoAndFechaUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetRegistroHabitoByHabitoAndFechaUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroHabitoByHabitoAndFechaUseCase(repo repositories.IRegistroHabito) *GetRegistroHabitoByHabitoAndFechaUseCase {
	return &GetRegistroHabitoByHabitoAndFechaUseCase{repo: repo}
}

func (uc *GetRegistroHabitoByHabitoAndFechaUseCase) Run(idHabito int32, fecha time.Time) (*entities.RegistroHabito, error) {
	return uc.repo.GetByHabitoAndFecha(idHabito, fecha)
}