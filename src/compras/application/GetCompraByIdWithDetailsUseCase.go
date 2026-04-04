// GetCompraByIdWithDetailsUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
)

type GetCompraByIdWithDetailsUseCase struct {
	repo repositories.ICompras
}

func NewGetCompraByIdWithDetailsUseCase(repo repositories.ICompras) *GetCompraByIdWithDetailsUseCase {
	return &GetCompraByIdWithDetailsUseCase{repo: repo}
}

func (uc *GetCompraByIdWithDetailsUseCase) Run(idCompra int32) (*repositories.CompraDetalles, error) {
	return uc.repo.GetCompraByIdWithDetails(idCompra)
}