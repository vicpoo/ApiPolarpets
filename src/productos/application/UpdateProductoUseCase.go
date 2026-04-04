// UpdateProductoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type UpdateProductoUseCase struct {
	repo repositories.IProductos
}

func NewUpdateProductoUseCase(repo repositories.IProductos) *UpdateProductoUseCase {
	return &UpdateProductoUseCase{repo: repo}
}

func (uc *UpdateProductoUseCase) Run(producto *entities.Productos) (*entities.Productos, error) {
	err := uc.repo.Update(producto)
	if err != nil {
		return nil, err
	}
	return producto, nil
}