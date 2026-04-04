// GetComprasRecientesByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetComprasRecientesByUserUseCase struct {
	repo repositories.ICompras
}

func NewGetComprasRecientesByUserUseCase(repo repositories.ICompras) *GetComprasRecientesByUserUseCase {
	return &GetComprasRecientesByUserUseCase{repo: repo}
}

func (uc *GetComprasRecientesByUserUseCase) Run(idUsuario int32, limit int) ([]entities.Compras, error) {
	return uc.repo.GetComprasRecientesByUser(idUsuario, limit)
}