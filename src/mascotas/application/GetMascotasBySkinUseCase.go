// GetMascotasBySkinUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetMascotasBySkinUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotasBySkinUseCase(repo repositories.IMascota) *GetMascotasBySkinUseCase {
	return &GetMascotasBySkinUseCase{repo: repo}
}

func (uc *GetMascotasBySkinUseCase) Run(idSkin int32) ([]entities.Mascota, error) {
	return uc.repo.GetBySkin(idSkin)
}