// GetComprasByProductoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetComprasByProductoUseCase struct {
	repo repositories.ICompras
}

func NewGetComprasByProductoUseCase(repo repositories.ICompras) *GetComprasByProductoUseCase {
	return &GetComprasByProductoUseCase{repo: repo}
}

func (uc *GetComprasByProductoUseCase) Run(idProducto int32) ([]entities.Compras, error) {
	return uc.repo.GetByProducto(idProducto)
}