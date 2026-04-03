// LoginUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type LoginUseCase struct {
	repo repositories.IUsuario
}

func NewLoginUseCase(repo repositories.IUsuario) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Run(email, password string) (*entities.Usuario, error) {
	return uc.repo.Login(email, password)
}