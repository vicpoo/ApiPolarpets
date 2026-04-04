// UpdateHabitoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type UpdateHabitoUseCase struct {
	repo repositories.IHabito
}

func NewUpdateHabitoUseCase(repo repositories.IHabito) *UpdateHabitoUseCase {
	return &UpdateHabitoUseCase{repo: repo}
}

func (uc *UpdateHabitoUseCase) Run(habito *entities.Habito) (*entities.Habito, error) {
	err := uc.repo.Update(habito)
	if err != nil {
		return nil, err
	}
	return habito, nil
}