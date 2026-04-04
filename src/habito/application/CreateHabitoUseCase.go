// CreateHabitoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type CreateHabitoUseCase struct {
	repo repositories.IHabito
}

func NewCreateHabitoUseCase(repo repositories.IHabito) *CreateHabitoUseCase {
	return &CreateHabitoUseCase{repo: repo}
}

func (uc *CreateHabitoUseCase) Run(habito *entities.Habito) (*entities.Habito, error) {
	err := uc.repo.Save(habito)
	if err != nil {
		return nil, err
	}
	return habito, nil
}