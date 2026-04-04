// GetRetoByTituloUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type GetRetoByTituloUseCase struct {
	repo repositories.IRetos
}

func NewGetRetoByTituloUseCase(repo repositories.IRetos) *GetRetoByTituloUseCase {
	return &GetRetoByTituloUseCase{repo: repo}
}

func (uc *GetRetoByTituloUseCase) Run(titulo string) (*entities.Retos, error) {
	return uc.repo.GetByTitulo(titulo)
}