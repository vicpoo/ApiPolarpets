// GetInventarioByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type GetInventarioByUserUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewGetInventarioByUserUseCase(repo repositories.IInventarioUsuario) *GetInventarioByUserUseCase {
	return &GetInventarioByUserUseCase{repo: repo}
}

func (uc *GetInventarioByUserUseCase) Run(idUsuario int32) ([]entities.InventarioUsuario, error) {
	return uc.repo.GetByUser(idUsuario)
}