// GetPendingRetosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetPendingRetosByUserUseCase struct {
	repo repositories.IUserRetos
}

func NewGetPendingRetosByUserUseCase(repo repositories.IUserRetos) *GetPendingRetosByUserUseCase {
	return &GetPendingRetosByUserUseCase{repo: repo}
}

func (uc *GetPendingRetosByUserUseCase) Run(idUsuario int32) ([]entities.UserRetos, error) {
	return uc.repo.GetPendingByUser(idUsuario)
}