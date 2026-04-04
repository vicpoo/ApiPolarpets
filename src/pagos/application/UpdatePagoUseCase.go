// UpdatePagoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type UpdatePagoUseCase struct {
	repo repositories.IPagos
}

func NewUpdatePagoUseCase(repo repositories.IPagos) *UpdatePagoUseCase {
	return &UpdatePagoUseCase{repo: repo}
}

func (uc *UpdatePagoUseCase) Run(pago *entities.Pagos) (*entities.Pagos, error) {
	err := uc.repo.Update(pago)
	if err != nil {
		return nil, err
	}
	return pago, nil
}