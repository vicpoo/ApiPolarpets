// GetMascotaCompletaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
)

type GetMascotaCompletaUseCase struct {
	repo repositories.IMascota
}

func NewGetMascotaCompletaUseCase(repo repositories.IMascota) *GetMascotaCompletaUseCase {
	return &GetMascotaCompletaUseCase{repo: repo}
}

func (uc *GetMascotaCompletaUseCase) Run(idMascota int32) (*repositories.MascotaDetalles, error) {
	return uc.repo.GetMascotaCompleta(idMascota)
}