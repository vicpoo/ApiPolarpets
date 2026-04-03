// RegisterUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type RegisterUseCase struct {
	repo repositories.IUsuario
}

func NewRegisterUseCase(repo repositories.IUsuario) *RegisterUseCase {
	return &RegisterUseCase{repo: repo}
}

func (uc *RegisterUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	err := uc.repo.Register(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}