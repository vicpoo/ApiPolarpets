// UpdateInventarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type UpdateInventarioUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewUpdateInventarioUseCase(repo repositories.IInventarioUsuario) *UpdateInventarioUseCase {
	return &UpdateInventarioUseCase{repo: repo}
}

func (uc *UpdateInventarioUseCase) Run(inventario *entities.InventarioUsuario) (*entities.InventarioUsuario, error) {
	err := uc.repo.Update(inventario)
	if err != nil {
		return nil, err
	}
	return inventario, nil
}