// GetInventarioByProductoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetInventarioByProductoUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByProductoUseCase(repo repositories.IInventarioUsuario) *GetInventarioByProductoUseCase {
	return &GetInventarioByProductoUseCase{repo: repo}
}

func (uc *GetInventarioByProductoUseCase) Run(idProducto int32) ([]entities.InventarioUsuario, error) {
	return uc.repo.GetByProducto(idProducto)
}