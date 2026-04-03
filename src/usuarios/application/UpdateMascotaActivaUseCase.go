// UpdateMascotaActivaUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"

type UpdateMascotaActivaUseCase struct {
	repo repositories.IUsuario
}

func NewUpdateMascotaActivaUseCase(repo repositories.IUsuario) *UpdateMascotaActivaUseCase {
	return &UpdateMascotaActivaUseCase{repo: repo}
}

func (uc *UpdateMascotaActivaUseCase) Run(idUsuario int32, idMascotaActiva *int32) error {
	return uc.repo.UpdateMascotaActiva(idUsuario, idMascotaActiva)
}