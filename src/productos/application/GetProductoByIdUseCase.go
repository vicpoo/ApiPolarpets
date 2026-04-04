// GetProductoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductoByIdUseCase struct {
	repo repositories.IProductos
}

func NewGetProductoByIdUseCase(repo repositories.IProductos) *GetProductoByIdUseCase {
	return &GetProductoByIdUseCase{repo: repo}
}

func (uc *GetProductoByIdUseCase) Run(id int32) (*entities.Productos, error) {
	return uc.repo.GetById(id)
}