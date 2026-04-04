// GetNivelByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type GetNivelByIdUseCase struct {
	repo repositories.INiveles
}

func NewGetNivelByIdUseCase(repo repositories.INiveles) *GetNivelByIdUseCase {
	return &GetNivelByIdUseCase{repo: repo}
}

func (uc *GetNivelByIdUseCase) Run(id int32) (*entities.Niveles, error) {
	return uc.repo.GetById(id)
}