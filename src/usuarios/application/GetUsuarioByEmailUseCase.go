// GetUsuarioByEmailUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type GetUsuarioByEmailUseCase struct {
	repo repositories.IUsuario
}

func NewGetUsuarioByEmailUseCase(repo repositories.IUsuario) *GetUsuarioByEmailUseCase {
	return &GetUsuarioByEmailUseCase{repo: repo}
}

func (uc *GetUsuarioByEmailUseCase) Run(email string) (*entities.Usuario, error) {
	return uc.repo.GetByEmail(email)
}