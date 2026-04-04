// GetComprasByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetComprasByUserUseCase struct {
	repo repositories.ICompras
}

func NewGetComprasByUserUseCase(repo repositories.ICompras) *GetComprasByUserUseCase {
	return &GetComprasByUserUseCase{repo: repo}
}

func (uc *GetComprasByUserUseCase) Run(idUsuario int32) ([]entities.Compras, error) {
	return uc.repo.GetByUser(idUsuario)
}