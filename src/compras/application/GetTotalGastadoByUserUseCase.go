// GetTotalGastadoByUserUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"

type GetTotalGastadoByUserUseCase struct {
	repo repositories.ICompras
}

func NewGetTotalGastadoByUserUseCase(repo repositories.ICompras) *GetTotalGastadoByUserUseCase {
	return &GetTotalGastadoByUserUseCase{repo: repo}
}

func (uc *GetTotalGastadoByUserUseCase) Run(idUsuario int32) (float64, error) {
	return uc.repo.GetTotalGastadoByUser(idUsuario)
}