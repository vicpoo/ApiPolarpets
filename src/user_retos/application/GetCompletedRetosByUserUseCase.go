// GetCompletedRetosByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type GetCompletedRetosByUserUseCase struct {
	repo repositories.IUserRetos
}

func NewGetCompletedRetosByUserUseCase(repo repositories.IUserRetos) *GetCompletedRetosByUserUseCase {
	return &GetCompletedRetosByUserUseCase{repo: repo}
}

func (uc *GetCompletedRetosByUserUseCase) Run(idUsuario int32) ([]entities.UserRetos, error) {
	return uc.repo.GetCompletedByUser(idUsuario)
}