// RegisterUseCase.go
package application

import (
	"errors"
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
	// Verificar si el email ya existe
	existingUser, _ := uc.repo.GetByEmail(usuario.GetEmail())
	if existingUser != nil {
		return nil, errors.New("el email ya está registrado")
	}

	// Hashear la contraseña
	err := usuario.HashPassword()
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	err = uc.repo.Register(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}