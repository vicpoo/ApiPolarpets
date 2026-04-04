// GetInventarioByUserAndProductWithDetailsUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
)

type GetInventarioByUserAndProductWithDetailsUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByUserAndProductWithDetailsUseCase(repo repositories.IInventarioUsuario) *GetInventarioByUserAndProductWithDetailsUseCase {
	return &GetInventarioByUserAndProductWithDetailsUseCase{repo: repo}
}

func (uc *GetInventarioByUserAndProductWithDetailsUseCase) Run(idUsuario, idProducto int32) (*repositories.InventarioDetalles, error) {
	return uc.repo.GetInventarioByUserAndProductWithDetails(idUsuario, idProducto)
}