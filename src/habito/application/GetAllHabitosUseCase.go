// GetAllHabitosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type GetAllHabitosUseCase struct {
	repo repositories.IHabito
}

func NewGetAllHabitosUseCase(repo repositories.IHabito) *GetAllHabitosUseCase {
	return &GetAllHabitosUseCase{repo: repo}
}

func (uc *GetAllHabitosUseCase) Run() ([]entities.Habito, error) {
	return uc.repo.GetAll()
}