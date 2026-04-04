// GetAllPagosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type GetAllPagosUseCase struct {
	repo repositories.IPagos
}

func NewGetAllPagosUseCase(repo repositories.IPagos) *GetAllPagosUseCase {
	return &GetAllPagosUseCase{repo: repo}
}

func (uc *GetAllPagosUseCase) Run() ([]entities.Pagos, error) {
	return uc.repo.GetAll()
}