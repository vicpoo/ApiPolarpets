// GetHabitoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type GetHabitoByIdUseCase struct {
	repo repositories.IHabito
}

func NewGetHabitoByIdUseCase(repo repositories.IHabito) *GetHabitoByIdUseCase {
	return &GetHabitoByIdUseCase{repo: repo}
}

func (uc *GetHabitoByIdUseCase) Run(id int32) (*entities.Habito, error) {
	return uc.repo.GetById(id)
}