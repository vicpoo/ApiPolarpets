// GetAllUsuariosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type GetAllUsuariosUseCase struct {
	repo repositories.IUsuario
}

func NewGetAllUsuariosUseCase(repo repositories.IUsuario) *GetAllUsuariosUseCase {
	return &GetAllUsuariosUseCase{repo: repo}
}

func (uc *GetAllUsuariosUseCase) Run() ([]entities.Usuario, error) {
	return uc.repo.GetAll()
}	