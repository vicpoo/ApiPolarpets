// GetUsuarioByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type GetUsuarioByIdUseCase struct {
	repo repositories.IUsuario
}

func NewGetUsuarioByIdUseCase(repo repositories.IUsuario) *GetUsuarioByIdUseCase {
	return &GetUsuarioByIdUseCase{repo: repo}
}

func (uc *GetUsuarioByIdUseCase) Run(id int32) (*entities.Usuario, error) {
	return uc.repo.GetById(id)
}