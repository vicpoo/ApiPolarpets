// GetPagoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagoByIdUseCase struct {
	repo repositories.IPagos
}

func NewGetPagoByIdUseCase(repo repositories.IPagos) *GetPagoByIdUseCase {
	return &GetPagoByIdUseCase{repo: repo}
}

func (uc *GetPagoByIdUseCase) Run(id int32) (*entities.Pagos, error) {
	return uc.repo.GetById(id)
}