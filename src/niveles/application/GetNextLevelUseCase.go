// GetNextLevelUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type GetNextLevelUseCase struct {
	repo repositories.INiveles
}

func NewGetNextLevelUseCase(repo repositories.INiveles) *GetNextLevelUseCase {
	return &GetNextLevelUseCase{repo: repo}
}

func (uc *GetNextLevelUseCase) Run(currentExp int32) (*entities.Niveles, error) {
	return uc.repo.GetNextLevel(currentExp)
}