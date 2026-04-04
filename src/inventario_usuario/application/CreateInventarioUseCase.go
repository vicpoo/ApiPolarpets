// CreateInventarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type CreateInventarioUseCase struct {
	repo repositories.IInventarioUsuario
}

func NewCreateInventarioUseCase(repo repositories.IInventarioUsuario) *CreateInventarioUseCase {
	return &CreateInventarioUseCase{repo: repo}
}

func (uc *CreateInventarioUseCase) Run(inventario *entities.InventarioUsuario) (*entities.InventarioUsuario, error) {
	err := uc.repo.Save(inventario)
	if err != nil {
		return nil, err
	}
	return inventario, nil
}