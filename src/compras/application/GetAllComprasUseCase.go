// GetAllComprasUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type GetAllComprasUseCase struct {
	repo repositories.ICompras
}

func NewGetAllComprasUseCase(repo repositories.ICompras) *GetAllComprasUseCase {
	return &GetAllComprasUseCase{repo: repo}
}

func (uc *GetAllComprasUseCase) Run() ([]entities.Compras, error) {
	return uc.repo.GetAll()
}