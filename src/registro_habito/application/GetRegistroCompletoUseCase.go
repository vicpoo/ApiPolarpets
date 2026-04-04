// GetRegistroCompletoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
)

type GetRegistroCompletoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetRegistroCompletoUseCase(repo repositories.IRegistroHabito) *GetRegistroCompletoUseCase {
	return &GetRegistroCompletoUseCase{repo: repo}
}

func (uc *GetRegistroCompletoUseCase) Run(idRegistro int32) (*repositories.RegistroHabitoDetalles, error) {
	return uc.repo.GetRegistroCompleto(idRegistro)
}