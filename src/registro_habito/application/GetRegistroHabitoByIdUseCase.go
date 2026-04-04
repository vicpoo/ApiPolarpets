// GetRegistroHabitoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetRegistroHabitoByIdUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroHabitoByIdUseCase(repo repositories.IRegistroHabito) *GetRegistroHabitoByIdUseCase {
	return &GetRegistroHabitoByIdUseCase{repo: repo}
}

func (uc *GetRegistroHabitoByIdUseCase) Run(id int32) (*entities.RegistroHabito, error) {
	return uc.repo.GetById(id)
}