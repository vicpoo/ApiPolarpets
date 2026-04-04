// GetInventarioByUserAndProductUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetInventarioByUserAndProductUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByUserAndProductUseCase(repo repositories.IInventarioUsuario) *GetInventarioByUserAndProductUseCase {
	return &GetInventarioByUserAndProductUseCase{repo: repo}
}

func (uc *GetInventarioByUserAndProductUseCase) Run(idUsuario, idProducto int32) (*entities.InventarioUsuario, error) {
	return uc.repo.GetByUserAndProduct(idUsuario, idProducto)
}