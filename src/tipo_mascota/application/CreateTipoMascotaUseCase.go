// CreateTipoMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type CreateTipoMascotaUseCase struct {
	repo repositories.ITipoMascota
}

func NewCreateTipoMascotaUseCase(repo repositories.ITipoMascota) *CreateTipoMascotaUseCase {
	return &CreateTipoMascotaUseCase{repo: repo}
}

func (uc *CreateTipoMascotaUseCase) Run(tipoMascota *entities.TipoMascota) (*entities.TipoMascota, error) {
	err := uc.repo.Save(tipoMascota)
	if err != nil {
		return nil, err
	}
	return tipoMascota, nil
}