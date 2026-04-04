// CreatePagoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type CreatePagoUseCase struct {
	repo repositories.IPagos
}

func NewCreatePagoUseCase(repo repositories.IPagos) *CreatePagoUseCase {
	return &CreatePagoUseCase{repo: repo}
}

func (uc *CreatePagoUseCase) Run(pago *entities.Pagos) (*entities.Pagos, error) {
	err := uc.repo.Save(pago)
	if err != nil {
		return nil, err
	}
	return pago, nil
}