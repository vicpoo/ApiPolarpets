// CreateProductoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type CreateProductoUseCase struct {
	repo repositories.IProductos
}

func NewCreateProductoUseCase(repo repositories.IProductos) *CreateProductoUseCase {
	return &CreateProductoUseCase{repo: repo}
}

func (uc *CreateProductoUseCase) Run(producto *entities.Productos) (*entities.Productos, error) {
	err := uc.repo.Save(producto)
	if err != nil {
		return nil, err
	}
	return producto, nil
}