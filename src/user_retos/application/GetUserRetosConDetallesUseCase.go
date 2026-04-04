// GetUserRetosConDetallesUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
)

type GetUserRetosConDetallesUseCase struct {
	repo repositories.IUserRetos
}

func NewGetUserRetosConDetallesUseCase(repo repositories.IUserRetos) *GetUserRetosConDetallesUseCase {
	return &GetUserRetosConDetallesUseCase{repo: repo}
}

func (uc *GetUserRetosConDetallesUseCase) Run(idUsuario int32) ([]repositories.UserRetoDetalles, error) {
	return uc.repo.GetUserRetosConDetalles(idUsuario)
}