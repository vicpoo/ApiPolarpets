// UpdateMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type UpdateMascotaUseCase struct {
	repo repositories.IMascota
}

func NewUpdateMascotaUseCase(repo repositories.IMascota) *UpdateMascotaUseCase {
	return &UpdateMascotaUseCase{repo: repo}
}

func (uc *UpdateMascotaUseCase) Run(mascota *entities.Mascota) (*entities.Mascota, error) {
	err := uc.repo.Update(mascota)
	if err != nil {
		return nil, err
	}
	return mascota, nil
}