// CreateRegistroHabitoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type CreateRegistroHabitoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewCreateRegistroHabitoUseCase(repo repositories.IRegistroHabito) *CreateRegistroHabitoUseCase {
	return &CreateRegistroHabitoUseCase{repo: repo}
}

func (uc *CreateRegistroHabitoUseCase) Run(registro *entities.RegistroHabito) (*entities.RegistroHabito, error) {
	err := uc.repo.Save(registro)
	if err != nil {
		return nil, err
	}
	return registro, nil
}