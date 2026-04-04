// GetInventarioByUserWithDetailsUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
)

type GetInventarioByUserWithDetailsUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByUserWithDetailsUseCase(repo repositories.IInventarioUsuario) *GetInventarioByUserWithDetailsUseCase {
	return &GetInventarioByUserWithDetailsUseCase{repo: repo}
}

func (uc *GetInventarioByUserWithDetailsUseCase) Run(idUsuario int32) ([]repositories.InventarioDetalles, error) {
	return uc.repo.GetInventarioByUserWithDetails(idUsuario)
}