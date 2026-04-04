// GetRetosByPuntosRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type GetRetosByPuntosRangeUseCase struct {
	repo repositories.IRetos
}

func NewGetRetosByPuntosRangeUseCase(repo repositories.IRetos) *GetRetosByPuntosRangeUseCase {
	return &GetRetosByPuntosRangeUseCase{repo: repo}
}

func (uc *GetRetosByPuntosRangeUseCase) Run(minPuntos, maxPuntos int32) ([]entities.Retos, error) {
	return uc.repo.GetByPuntosRange(minPuntos, maxPuntos)
}