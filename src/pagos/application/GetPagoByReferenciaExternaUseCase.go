// GetPagoByReferenciaExternaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetPagoByReferenciaExternaUseCase struct {
	repo repositories.IPagos
}

func NewGetPagoByReferenciaExternaUseCase(repo repositories.IPagos) *GetPagoByReferenciaExternaUseCase {
	return &GetPagoByReferenciaExternaUseCase{repo: repo}
}

func (uc *GetPagoByReferenciaExternaUseCase) Run(referencia string) (*entities.Pagos, error) {
	return uc.repo.GetByReferenciaExterna(referencia)
}