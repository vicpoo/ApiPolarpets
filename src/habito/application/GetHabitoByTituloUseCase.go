// GetHabitoByTituloUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type GetHabitoByTituloUseCase struct {
	repo repositories.IHabito
}

func NewGetHabitoByTituloUseCase(repo repositories.IHabito) *GetHabitoByTituloUseCase {
	return &GetHabitoByTituloUseCase{repo: repo}
}

func (uc *GetHabitoByTituloUseCase) Run(titulo string) (*entities.Habito, error) {
	return uc.repo.GetByTitulo(titulo)
}