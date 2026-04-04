// GetNivelByExpRequeridaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type GetNivelByExpRequeridaUseCase struct {
	repo repositories.INiveles
}

func NewGetNivelByExpRequeridaUseCase(repo repositories.INiveles) *GetNivelByExpRequeridaUseCase {
	return &GetNivelByExpRequeridaUseCase{repo: repo}
}

func (uc *GetNivelByExpRequeridaUseCase) Run(expRequerida int32) (*entities.Niveles, error) {
	return uc.repo.GetByExpRequerida(expRequerida)
}