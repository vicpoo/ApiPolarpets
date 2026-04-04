// GetUserRetoByUserAndRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetUserRetoByUserAndRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewGetUserRetoByUserAndRetoUseCase(repo repositories.IUserRetos) *GetUserRetoByUserAndRetoUseCase {
	return &GetUserRetoByUserAndRetoUseCase{repo: repo}
}

func (uc *GetUserRetoByUserAndRetoUseCase) Run(idUsuario, idReto int32) (*entities.UserRetos, error) {
	return uc.repo.GetByUserAndReto(idUsuario, idReto)
}