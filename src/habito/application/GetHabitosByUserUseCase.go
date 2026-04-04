// GetHabitosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type GetHabitosByUserUseCase struct {
	repo repositories.IHabito
}

func NewGetHabitosByUserUseCase(repo repositories.IHabito) *GetHabitosByUserUseCase {
	return &GetHabitosByUserUseCase{repo: repo}
}

func (uc *GetHabitosByUserUseCase) Run(idUser int32) ([]entities.Habito, error) {
	return uc.repo.GetByUser(idUser)
}