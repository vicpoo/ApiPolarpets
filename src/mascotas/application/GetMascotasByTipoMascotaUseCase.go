// GetMascotasByTipoMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetMascotasByTipoMascotaUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotasByTipoMascotaUseCase(repo repositories.IMascota) *GetMascotasByTipoMascotaUseCase {
	return &GetMascotasByTipoMascotaUseCase{repo: repo}
}

func (uc *GetMascotasByTipoMascotaUseCase) Run(idTipoMascota int32) ([]entities.Mascota, error) {
	return uc.repo.GetByTipoMascota(idTipoMascota)
}