// UpdateRegistroHabitoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type UpdateRegistroHabitoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewUpdateRegistroHabitoUseCase(repo repositories.IRegistroHabito) *UpdateRegistroHabitoUseCase {
	return &UpdateRegistroHabitoUseCase{repo: repo}
}

func (uc *UpdateRegistroHabitoUseCase) Run(registro *entities.RegistroHabito) (*entities.RegistroHabito, error) {
	err := uc.repo.Update(registro)
	if err != nil {
		return nil, err
	}
	return registro, nil
}