// GetUserRetoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetUserRetoByIdUseCase struct {
	repo repositories.IUserRetos
}

func NewGetUserRetoByIdUseCase(repo repositories.IUserRetos) *GetUserRetoByIdUseCase {
	return &GetUserRetoByIdUseCase{repo: repo}
}

func (uc *GetUserRetoByIdUseCase) Run(id int32) (*entities.UserRetos, error) {
	return uc.repo.GetById(id)
}