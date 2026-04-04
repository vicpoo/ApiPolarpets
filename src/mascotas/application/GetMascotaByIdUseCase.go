// GetMascotaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetMascotaByIdUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotaByIdUseCase(repo repositories.IMascota) *GetMascotaByIdUseCase {
	return &GetMascotaByIdUseCase{repo: repo}
}

func (uc *GetMascotaByIdUseCase) Run(id int32) (*entities.Mascota, error) {
	return uc.repo.GetById(id)
}