// DeleteInventarioByUserAndProductUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"

type DeleteInventarioByUserAndProductUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewDeleteInventarioByUserAndProductUseCase(repo repositories.IInventarioUsuario) *DeleteInventarioByUserAndProductUseCase {
	return &DeleteInventarioByUserAndProductUseCase{repo: repo}
}

func (uc *DeleteInventarioByUserAndProductUseCase) Run(idUsuario, idProducto int32) error {
	return uc.repo.DeleteByUserAndProduct(idUsuario, idProducto)
}