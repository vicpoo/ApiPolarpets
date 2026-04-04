// GetProductoByNombreUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductoByNombreUseCase struct {
	repo repositories.IProductos
}

func NewGetProductoByNombreUseCase(repo repositories.IProductos) *GetProductoByNombreUseCase {
	return &GetProductoByNombreUseCase{repo: repo}
}

func (uc *GetProductoByNombreUseCase) Run(nombre string) (*entities.Productos, error) {
	return uc.repo.GetByNombre(nombre)
}