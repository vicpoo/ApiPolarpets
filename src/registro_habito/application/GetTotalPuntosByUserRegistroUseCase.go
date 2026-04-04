// GetTotalPuntosByUserRegistroUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"

type GetTotalPuntosByUserRegistroUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetTotalPuntosByUserRegistroUseCase(repo repositories.IRegistroHabito) *GetTotalPuntosByUserRegistroUseCase {
	return &GetTotalPuntosByUserRegistroUseCase{repo: repo}
}

func (uc *GetTotalPuntosByUserRegistroUseCase) Run(idUser int32) (int32, error) {
	return uc.repo.GetTotalPuntosByUser(idUser)
}