// GetTipoMascotaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type GetTipoMascotaByIdUseCase struct {
	repo repositories.ITipoMascota
}

func NewGetTipoMascotaByIdUseCase(repo repositories.ITipoMascota) *GetTipoMascotaByIdUseCase {
	return &GetTipoMascotaByIdUseCase{repo: repo}
}

func (uc *GetTipoMascotaByIdUseCase) Run(id int32) (*entities.TipoMascota, error) {
	return uc.repo.GetById(id)
}