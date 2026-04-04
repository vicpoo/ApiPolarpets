// GetAllProductosConDetallesUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
)

type GetAllProductosConDetallesUseCase struct {
	repo repositories.IProductos
}

func NewGetAllProductosConDetallesUseCase(repo repositories.IProductos) *GetAllProductosConDetallesUseCase {
	return &GetAllProductosConDetallesUseCase{repo: repo}
}

func (uc *GetAllProductosConDetallesUseCase) Run() ([]repositories.ProductoDetalles, error) {
	return uc.repo.GetAllProductosConDetalles()
}