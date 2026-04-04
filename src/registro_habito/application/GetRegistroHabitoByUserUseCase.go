// GetRegistroHabitoByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type GetRegistroHabitoByUserUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroHabitoByUserUseCase(repo repositories.IRegistroHabito) *GetRegistroHabitoByUserUseCase {
	return &GetRegistroHabitoByUserUseCase{repo: repo}
}

func (uc *GetRegistroHabitoByUserUseCase) Run(idUser int32) ([]entities.RegistroHabito, error) {
	return uc.repo.GetByUser(idUser)
}