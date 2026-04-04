// GetNivelByNivelUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type GetNivelByNivelUseCase struct {
	repo repositories.INiveles
}

func NewGetNivelByNivelUseCase(repo repositories.INiveles) *GetNivelByNivelUseCase {
	return &GetNivelByNivelUseCase{repo: repo}
}

func (uc *GetNivelByNivelUseCase) Run(nivel int32) (*entities.Niveles, error) {
	return uc.repo.GetByNivel(nivel)
}