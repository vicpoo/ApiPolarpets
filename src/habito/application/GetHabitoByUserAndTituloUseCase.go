// GetHabitoByUserAndTituloUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type GetHabitoByUserAndTituloUseCase struct {
	repo repositories.IHabito
}

func NewGetHabitoByUserAndTituloUseCase(repo repositories.IHabito) *GetHabitoByUserAndTituloUseCase {
	return &GetHabitoByUserAndTituloUseCase{repo: repo}
}

func (uc *GetHabitoByUserAndTituloUseCase) Run(idUser int32, titulo string) (*entities.Habito, error) {
	return uc.repo.GetByUserAndTitulo(idUser, titulo)
}