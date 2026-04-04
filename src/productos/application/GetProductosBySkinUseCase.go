// GetProductosBySkinUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductosBySkinUseCase struct {
	repo repositories.IProductos
}

func NewGetProductosBySkinUseCase(repo repositories.IProductos) *GetProductosBySkinUseCase {
	return &GetProductosBySkinUseCase{repo: repo}
}

func (uc *GetProductosBySkinUseCase) Run(idSkin int32) ([]entities.Productos, error) {
	return uc.repo.GetBySkin(idSkin)
}