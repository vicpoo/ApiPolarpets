// GetProductosByTipoInInventoryUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetProductosByTipoInInventoryUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetProductosByTipoInInventoryUseCase(repo repositories.IInventarioUsuario) *GetProductosByTipoInInventoryUseCase {
	return &GetProductosByTipoInInventoryUseCase{repo: repo}
}

func (uc *GetProductosByTipoInInventoryUseCase) Run(idUsuario int32, tipo string) ([]entities.InventarioUsuario, error) {
	return uc.repo.GetProductosByTipoInInventory(idUsuario, tipo)
}