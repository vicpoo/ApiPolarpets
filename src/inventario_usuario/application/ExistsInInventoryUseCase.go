// ExistsInInventoryUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"

type ExistsInInventoryUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewExistsInInventoryUseCase(repo repositories.IInventarioUsuario) *ExistsInInventoryUseCase {
	return &ExistsInInventoryUseCase{repo: repo}
}

func (uc *ExistsInInventoryUseCase) Run(idUsuario, idProducto int32) (bool, error) {
	return uc.repo.ExistsInInventory(idUsuario, idProducto)
}