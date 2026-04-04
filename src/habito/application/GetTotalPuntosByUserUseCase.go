// GetTotalPuntosByUserUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"

type GetTotalPuntosByUserUseCase struct {
	repo repositories.IHabito
}

func NewGetTotalPuntosByUserUseCase(repo repositories.IHabito) *GetTotalPuntosByUserUseCase {
	return &GetTotalPuntosByUserUseCase{repo: repo}
}

func (uc *GetTotalPuntosByUserUseCase) Run(idUser int32) (int32, error) {
	return uc.repo.GetTotalPuntosByUser(idUser)
}