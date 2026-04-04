// GetSkinByTipoMascotaAndNombreUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type GetSkinByTipoMascotaAndNombreUseCase struct {
	repo repositories.ISkins
}

func NewGetSkinByTipoMascotaAndNombreUseCase(repo repositories.ISkins) *GetSkinByTipoMascotaAndNombreUseCase {
	return &GetSkinByTipoMascotaAndNombreUseCase{repo: repo}
}

func (uc *GetSkinByTipoMascotaAndNombreUseCase) Run(idTipoMascota int32, nombre string) (*entities.Skins, error) {
	return uc.repo.GetByTipoMascotaAndNombre(idTipoMascota, nombre)
}