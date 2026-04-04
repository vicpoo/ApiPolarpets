// GetAllTiposMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type GetAllTiposMascotaUseCase struct {
	repo repositories.ITipoMascota
}

func NewGetAllTiposMascotaUseCase(repo repositories.ITipoMascota) *GetAllTiposMascotaUseCase {
	return &GetAllTiposMascotaUseCase{repo: repo}
}

func (uc *GetAllTiposMascotaUseCase) Run() ([]entities.TipoMascota, error) {
	return uc.repo.GetAll()
}