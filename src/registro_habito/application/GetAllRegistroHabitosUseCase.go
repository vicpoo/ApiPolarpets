// GetAllRegistroHabitosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetAllRegistroHabitosUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetAllRegistroHabitosUseCase(repo repositories.IRegistroHabito) *GetAllRegistroHabitosUseCase {
	return &GetAllRegistroHabitosUseCase{repo: repo}
}

func (uc *GetAllRegistroHabitosUseCase) Run() ([]entities.RegistroHabito, error) {
	return uc.repo.GetAll()
}