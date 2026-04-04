// GetProductosByPrecioRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductosByPrecioRangeUseCase struct {
	repo repositories.IProductos
}

func NewGetProductosByPrecioRangeUseCase(repo repositories.IProductos) *GetProductosByPrecioRangeUseCase {
	return &GetProductosByPrecioRangeUseCase{repo: repo}
}

func (uc *GetProductosByPrecioRangeUseCase) Run(minPrecio, maxPrecio float64) ([]entities.Productos, error) {
	return uc.repo.GetByPrecioRange(minPrecio, maxPrecio)
}