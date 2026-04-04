// DeleteProductoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"

type DeleteProductoUseCase struct {
	repo repositories.IProductos
}

func NewDeleteProductoUseCase(repo repositories.IProductos) *DeleteProductoUseCase {
	return &DeleteProductoUseCase{repo: repo}
}

func (uc *DeleteProductoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}