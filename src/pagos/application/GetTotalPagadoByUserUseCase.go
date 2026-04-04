// GetTotalPagadoByUserUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"

type GetTotalPagadoByUserUseCase struct {
	repo repositories.IPagos
}

func NewGetTotalPagadoByUserUseCase(repo repositories.IPagos) *GetTotalPagadoByUserUseCase {
	return &GetTotalPagadoByUserUseCase{repo: repo}
}

func (uc *GetTotalPagadoByUserUseCase) Run(idUsuario int32) (float64, error) {
	return uc.repo.GetTotalPagadoByUser(idUsuario)
}