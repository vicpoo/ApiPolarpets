// GetUserRetosByRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetUserRetosByRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewGetUserRetosByRetoUseCase(repo repositories.IUserRetos) *GetUserRetosByRetoUseCase {
	return &GetUserRetosByRetoUseCase{repo: repo}
}

func (uc *GetUserRetosByRetoUseCase) Run(idReto int32) ([]entities.UserRetos, error) {
	return uc.repo.GetByReto(idReto)
}