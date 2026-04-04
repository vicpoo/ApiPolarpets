// GetProductosByTipoMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type GetProductosByTipoMascotaUseCase struct {
	repo repositories.IProductos
}

func NewGetProductosByTipoMascotaUseCase(repo repositories.IProductos) *GetProductosByTipoMascotaUseCase {
	return &GetProductosByTipoMascotaUseCase{repo: repo}
}

func (uc *GetProductosByTipoMascotaUseCase) Run(idTipoMascota int32) ([]entities.Productos, error) {
	return uc.repo.GetByTipoMascota(idTipoMascota)
}