// GetCompraByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetCompraByIdUseCase struct {
	repo repositories.ICompras
}

func NewGetCompraByIdUseCase(repo repositories.ICompras) *GetCompraByIdUseCase {
	return &GetCompraByIdUseCase{repo: repo}
}

func (uc *GetCompraByIdUseCase) Run(id int32) (*entities.Compras, error) {
	return uc.repo.GetById(id)
}