// GetAllProductosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetAllProductosUseCase struct {
	repo repositories.IProductos
}

func NewGetAllProductosUseCase(repo repositories.IProductos) *GetAllProductosUseCase {
	return &GetAllProductosUseCase{repo: repo}
}

func (uc *GetAllProductosUseCase) Run() ([]entities.Productos, error) {
	return uc.repo.GetAll()
}