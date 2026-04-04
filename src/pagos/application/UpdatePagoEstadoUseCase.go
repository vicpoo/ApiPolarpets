// UpdatePagoEstadoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"

type UpdatePagoEstadoUseCase struct {
	repo repositories.IPagos
}

func NewUpdatePagoEstadoUseCase(repo repositories.IPagos) *UpdatePagoEstadoUseCase {
	return &UpdatePagoEstadoUseCase{repo: repo}
}

func (uc *UpdatePagoEstadoUseCase) Run(id int32, estado string) error {
	return uc.repo.UpdateEstado(id, estado)
}