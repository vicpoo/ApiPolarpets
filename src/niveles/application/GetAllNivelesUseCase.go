// GetAllNivelesUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type GetAllNivelesUseCase struct {
	repo repositories.INiveles
}

func NewGetAllNivelesUseCase(repo repositories.INiveles) *GetAllNivelesUseCase {
	return &GetAllNivelesUseCase{repo: repo}
}

func (uc *GetAllNivelesUseCase) Run() ([]entities.Niveles, error) {
	return uc.repo.GetAll()
}