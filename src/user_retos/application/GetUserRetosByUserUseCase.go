// GetUserRetosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetUserRetosByUserUseCase struct {
	repo repositories.IUserRetos
}

func NewGetUserRetosByUserUseCase(repo repositories.IUserRetos) *GetUserRetosByUserUseCase {
	return &GetUserRetosByUserUseCase{repo: repo}
}

func (uc *GetUserRetosByUserUseCase) Run(idUsuario int32) ([]entities.UserRetos, error) {
	return uc.repo.GetByUser(idUsuario)
}