// DeleteCompraUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"

type DeleteCompraUseCase struct {
	repo repositories.ICompras
}

func NewDeleteCompraUseCase(repo repositories.ICompras) *DeleteCompraUseCase {
	return &DeleteCompraUseCase{repo: repo}
}

func (uc *DeleteCompraUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}