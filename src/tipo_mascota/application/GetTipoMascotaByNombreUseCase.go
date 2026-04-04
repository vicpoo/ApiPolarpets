// GetTipoMascotaByNombreUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type GetTipoMascotaByNombreUseCase struct {
	repo repositories.ITipoMascota
}

func NewGetTipoMascotaByNombreUseCase(repo repositories.ITipoMascota) *GetTipoMascotaByNombreUseCase {
	return &GetTipoMascotaByNombreUseCase{repo: repo}
}

func (uc *GetTipoMascotaByNombreUseCase) Run(nombre string) (*entities.TipoMascota, error) {
	return uc.repo.GetByNombre(nombre)
}