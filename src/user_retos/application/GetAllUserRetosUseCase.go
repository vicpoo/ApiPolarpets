// GetAllUserRetosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetAllUserRetosUseCase struct {
	repo repositories.IUserRetos
}

func NewGetAllUserRetosUseCase(repo repositories.IUserRetos) *GetAllUserRetosUseCase {
	return &GetAllUserRetosUseCase{repo: repo}
}

func (uc *GetAllUserRetosUseCase) Run() ([]entities.UserRetos, error) {
	return uc.repo.GetAll()
}