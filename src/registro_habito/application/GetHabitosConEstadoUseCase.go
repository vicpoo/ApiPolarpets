// GetHabitosConEstadoUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetHabitosConEstadoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetHabitosConEstadoUseCase(repo repositories.IRegistroHabito) *GetHabitosConEstadoUseCase {
	return &GetHabitosConEstadoUseCase{repo: repo}
}

func (uc *GetHabitosConEstadoUseCase) Run(idUser int32, fecha time.Time) ([]entities.HabitoConEstado, error) {
	return uc.repo.GetHabitosConEstadoByFecha(idUser, fecha)
}