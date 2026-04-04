// DeleteTipoMascotaUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"

type DeleteTipoMascotaUseCase struct {
	repo repositories.ITipoMascota
}

func NewDeleteTipoMascotaUseCase(repo repositories.ITipoMascota) *DeleteTipoMascotaUseCase {
	return &DeleteTipoMascotaUseCase{repo: repo}
}

func (uc *DeleteTipoMascotaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}