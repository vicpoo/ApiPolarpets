// GetMascotasByNivelUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetMascotasByNivelUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotasByNivelUseCase(repo repositories.IMascota) *GetMascotasByNivelUseCase {
	return &GetMascotasByNivelUseCase{repo: repo}
}

func (uc *GetMascotasByNivelUseCase) Run(idNiveles int32) ([]entities.Mascota, error) {
	return uc.repo.GetByNivel(idNiveles)
}