// GetAllRetosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type GetAllRetosUseCase struct {
	repo repositories.IRetos
}

func NewGetAllRetosUseCase(repo repositories.IRetos) *GetAllRetosUseCase {
	return &GetAllRetosUseCase{repo: repo}
}

func (uc *GetAllRetosUseCase) Run() ([]entities.Retos, error) {
	return uc.repo.GetAll()
}