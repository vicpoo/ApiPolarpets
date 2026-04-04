// GetMascotasByUserUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetMascotasByUserUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotasByUserUseCase(repo repositories.IMascota) *GetMascotasByUserUseCase {
	return &GetMascotasByUserUseCase{repo: repo}
}

func (uc *GetMascotasByUserUseCase) Run(idUser int32) ([]entities.Mascota, error) {
	return uc.repo.GetByUser(idUser)
}