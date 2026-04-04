// GetProductoConDetallesUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
)

type GetProductoConDetallesUseCase struct {
	repo repositories.IProductos
}

func NewGetProductoConDetallesUseCase(repo repositories.IProductos) *GetProductoConDetallesUseCase {
	return &GetProductoConDetallesUseCase{repo: repo}
}

func (uc *GetProductoConDetallesUseCase) Run(id int32) (*repositories.ProductoDetalles, error) {
	return uc.repo.GetProductosConDetalles(id)
}