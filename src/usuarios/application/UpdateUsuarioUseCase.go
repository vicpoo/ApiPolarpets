// UpdateUsuarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type UpdateUsuarioUseCase struct {
	repo repositories.IUsuario
}

func NewUpdateUsuarioUseCase(repo repositories.IUsuario) *UpdateUsuarioUseCase {
	return &UpdateUsuarioUseCase{repo: repo}
}

func (uc *UpdateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	// Obtener usuario actual para verificar si la contraseña cambió
	existingUser, err := uc.repo.GetById(usuario.GetIDUsuario())
	if err != nil {
		return nil, err
	}

	// Si la contraseña es diferente a la almacenada, hashearla
	if usuario.GetPassword() != existingUser.GetPassword() {
		err := usuario.HashPassword()
		if err != nil {
			return nil, err
		}
	}
	
	err = uc.repo.Update(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}