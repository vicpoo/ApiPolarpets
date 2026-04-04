// GetCompraByPagoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetCompraByPagoUseCase struct {
	repo repositories.ICompras
}

func NewGetCompraByPagoUseCase(repo repositories.ICompras) *GetCompraByPagoUseCase {
	return &GetCompraByPagoUseCase{repo: repo}
}

func (uc *GetCompraByPagoUseCase) Run(idPago int32) (*entities.Compras, error) {
	return uc.repo.GetByPago(idPago)
}