// GetAllMascotasUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type GetAllMascotasUseCase struct {
	repo repositories.IMascota
}

func NewGetAllMascotasUseCase(repo repositories.IMascota) *GetAllMascotasUseCase {
	return &GetAllMascotasUseCase{repo: repo}
}

func (uc *GetAllMascotasUseCase) Run() ([]entities.Mascota, error) {
	return uc.repo.GetAll()
}