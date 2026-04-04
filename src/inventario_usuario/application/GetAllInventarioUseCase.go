// GetAllInventarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetAllInventarioUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetAllInventarioUseCase(repo repositories.IInventarioUsuario) *GetAllInventarioUseCase {
	return &GetAllInventarioUseCase{repo: repo}
}

func (uc *GetAllInventarioUseCase) Run() ([]entities.InventarioUsuario, error) {
	return uc.repo.GetAll()
}