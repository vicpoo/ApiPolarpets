// CreateNivelUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type CreateNivelUseCase struct {
	repo repositories.INiveles
}

func NewCreateNivelUseCase(repo repositories.INiveles) *CreateNivelUseCase {
	return &CreateNivelUseCase{repo: repo}
}

func (uc *CreateNivelUseCase) Run(nivel *entities.Niveles) (*entities.Niveles, error) {
	err := uc.repo.Save(nivel)
	if err != nil {
		return nil, err
	}
	return nivel, nil
}