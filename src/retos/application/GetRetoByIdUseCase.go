// GetRetoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type GetRetoByIdUseCase struct {
	repo repositories.IRetos
}

func NewGetRetoByIdUseCase(repo repositories.IRetos) *GetRetoByIdUseCase {
	return &GetRetoByIdUseCase{repo: repo}
}

func (uc *GetRetoByIdUseCase) Run(id int32) (*entities.Retos, error) {
	return uc.repo.GetById(id)
}