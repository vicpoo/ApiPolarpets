// GetProductosByTipoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductosByTipoUseCase struct {
	repo repositories.IProductos
}

func NewGetProductosByTipoUseCase(repo repositories.IProductos) *GetProductosByTipoUseCase {
	return &GetProductosByTipoUseCase{repo: repo}
}

func (uc *GetProductosByTipoUseCase) Run(tipo string) ([]entities.Productos, error) {
	return uc.repo.GetByTipo(tipo)
}