// CreateCompraUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type CreateCompraUseCase struct {
	repo repositories.ICompras
}

func NewCreateCompraUseCase(repo repositories.ICompras) *CreateCompraUseCase {
	return &CreateCompraUseCase{repo: repo}
}

func (uc *CreateCompraUseCase) Run(compra *entities.Compras) (*entities.Compras, error) {
	err := uc.repo.Save(compra)
	if err != nil {
		return nil, err
	}
	return compra, nil
}