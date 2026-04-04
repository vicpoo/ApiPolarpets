// GetCantidadProductosByUserUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"

type GetCantidadProductosByUserUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetCantidadProductosByUserUseCase(repo repositories.IInventarioUsuario) *GetCantidadProductosByUserUseCase {
	return &GetCantidadProductosByUserUseCase{repo: repo}
}

func (uc *GetCantidadProductosByUserUseCase) Run(idUsuario int32) (int32, error) {
	return uc.repo.GetCantidadProductosByUser(idUsuario)
}