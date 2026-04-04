// GetSkinsByTipoMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type GetSkinsByTipoMascotaUseCase struct {
	repo repositories.ISkins
}

func NewGetSkinsByTipoMascotaUseCase(repo repositories.ISkins) *GetSkinsByTipoMascotaUseCase {
	return &GetSkinsByTipoMascotaUseCase{repo: repo}
}

func (uc *GetSkinsByTipoMascotaUseCase) Run(idTipoMascota int32) ([]entities.Skins, error) {
	return uc.repo.GetByTipoMascota(idTipoMascota)
}