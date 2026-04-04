// GetRegistroHabitoByHabitoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetRegistroHabitoByHabitoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroHabitoByHabitoUseCase(repo repositories.IRegistroHabito) *GetRegistroHabitoByHabitoUseCase {
	return &GetRegistroHabitoByHabitoUseCase{repo: repo}
}

func (uc *GetRegistroHabitoByHabitoUseCase) Run(idHabito int32) ([]entities.RegistroHabito, error) {
	return uc.repo.GetByHabito(idHabito)
}