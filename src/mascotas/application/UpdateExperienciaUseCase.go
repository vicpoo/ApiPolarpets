// UpdateExperienciaUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"

type UpdateExperienciaUseCase struct {
	repo repositories.IMascota
}

func NewUpdateExperienciaUseCase(repo repositories.IMascota) *UpdateExperienciaUseCase {
	return &UpdateExperienciaUseCase{repo: repo}
}

func (uc *UpdateExperienciaUseCase) Run(idMascota int32, nuevaExperiencia int32) error {
	return uc.repo.UpdateExperiencia(idMascota, nuevaExperiencia)
}