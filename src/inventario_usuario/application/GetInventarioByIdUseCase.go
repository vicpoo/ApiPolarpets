// GetInventarioByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetInventarioByIdUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByIdUseCase(repo repositories.IInventarioUsuario) *GetInventarioByIdUseCase {
	return &GetInventarioByIdUseCase{repo: repo}
}

func (uc *GetInventarioByIdUseCase) Run(id int32) (*entities.InventarioUsuario, error) {
	return uc.repo.GetById(id)
}