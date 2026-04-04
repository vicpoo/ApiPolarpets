// UpdateNivelUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type UpdateNivelUseCase struct {
	repo repositories.INiveles
}

func NewUpdateNivelUseCase(repo repositories.INiveles) *UpdateNivelUseCase {
	return &UpdateNivelUseCase{repo: repo}
}

func (uc *UpdateNivelUseCase) Run(nivel *entities.Niveles) (*entities.Niveles, error) {
	err := uc.repo.Update(nivel)
	if err != nil {
		return nil, err
	}
	return nivel, nil
}