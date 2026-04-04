// UpdateCompraUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type UpdateCompraUseCase struct {
	repo repositories.ICompras
}

func NewUpdateCompraUseCase(repo repositories.ICompras) *UpdateCompraUseCase {
	return &UpdateCompraUseCase{repo: repo}
}

func (uc *UpdateCompraUseCase) Run(compra *entities.Compras) (*entities.Compras, error) {
	err := uc.repo.Update(compra)
	if err != nil {
		return nil, err
	}
	return compra, nil
}