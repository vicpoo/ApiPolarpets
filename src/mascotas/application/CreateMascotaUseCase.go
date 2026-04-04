// CreateMascotaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type CreateMascotaUseCase struct {
	repo repositories.IMascota
}

func NewCreateMascotaUseCase(repo repositories.IMascota) *CreateMascotaUseCase {
	return &CreateMascotaUseCase{repo: repo}
}

func (uc *CreateMascotaUseCase) Run(mascota *entities.Mascota) (*entities.Mascota, error) {
	err := uc.repo.Save(mascota)
	if err != nil {
		return nil, err
	}
	return mascota, nil
}