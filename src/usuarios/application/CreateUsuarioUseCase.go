// CreateUsuarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type CreateUsuarioUseCase struct {
	repo repositories.IUsuario
}

func NewCreateUsuarioUseCase(repo repositories.IUsuario) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repo: repo}
}

func (uc *CreateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	// Hashear la contraseña
	err := usuario.HashPassword()
	if err != nil {
		return nil, err
	}
	
	err = uc.repo.Save(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}