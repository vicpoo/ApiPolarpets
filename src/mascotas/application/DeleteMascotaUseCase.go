// DeleteMascotaUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"

type DeleteMascotaUseCase struct {
	repo repositories.IMascota
}

func NewDeleteMascotaUseCase(repo repositories.IMascota) *DeleteMascotaUseCase {
	return &DeleteMascotaUseCase{repo: repo}
}

func (uc *DeleteMascotaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}