// DeleteInventarioUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"

type DeleteInventarioUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewDeleteInventarioUseCase(repo repositories.IInventarioUsuario) *DeleteInventarioUseCase {
	return &DeleteInventarioUseCase{repo: repo}
}

func (uc *DeleteInventarioUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}