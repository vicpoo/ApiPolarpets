// GetComprasByUserWithDetailsUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
)

type GetComprasByUserWithDetailsUseCase struct {
	repo repositories.ICompras
}

func NewGetComprasByUserWithDetailsUseCase(repo repositories.ICompras) *GetComprasByUserWithDetailsUseCase {
	return &GetComprasByUserWithDetailsUseCase{repo: repo}
}

func (uc *GetComprasByUserWithDetailsUseCase) Run(idUsuario int32) ([]repositories.CompraDetalles, error) {
	return uc.repo.GetComprasByUserWithDetails(idUsuario)
}