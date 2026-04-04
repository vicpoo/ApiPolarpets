// DeletePagoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"

type DeletePagoUseCase struct {
	repo repositories.IPagos
}

func NewDeletePagoUseCase(repo repositories.IPagos) *DeletePagoUseCase {
	return &DeletePagoUseCase{repo: repo}
}

func (uc *DeletePagoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}