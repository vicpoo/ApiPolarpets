// UpdateTipoMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type UpdateTipoMascotaUseCase struct {
	repo repositories.ITipoMascota
}

func NewUpdateTipoMascotaUseCase(repo repositories.ITipoMascota) *UpdateTipoMascotaUseCase {
	return &UpdateTipoMascotaUseCase{repo: repo}
}

func (uc *UpdateTipoMascotaUseCase) Run(tipoMascota *entities.TipoMascota) (*entities.TipoMascota, error) {
	err := uc.repo.Update(tipoMascota)
	if err != nil {
		return nil, err
	}
	return tipoMascota, nil
}