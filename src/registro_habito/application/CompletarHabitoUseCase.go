// CompletarHabitoUseCase.go
package application

import (
	"fmt"

	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	habitoRepositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
)

type CompletarHabitoUseCase struct {
	registroRepo repositories.IRegistroHabito
	habitoRepo   habitoRepositories.IHabito
}

func NewCompletarHabitoUseCase(
	registroRepo repositories.IRegistroHabito,
	habitoRepo habitoRepositories.IHabito,
) *CompletarHabitoUseCase {
	return &CompletarHabitoUseCase{
		registroRepo: registroRepo,
		habitoRepo:   habitoRepo,
	}
}

func (uc *CompletarHabitoUseCase) Run(idHabito int32, idUser int32) error {
	// Verificar si ya fue completado hoy
	exists, err := uc.registroRepo.ExistsRegistroHoy(idHabito)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("el hábito ya fue completado hoy")
	}

	// Obtener el hábito para saber los puntos
	habito, err := uc.habitoRepo.GetById(idHabito)
	if err != nil {
		return err
	}

	// Verificar que el hábito pertenezca al usuario
	if habito.GetIDUser() != idUser {
		return fmt.Errorf("el hábito no pertenece al usuario")
	}

	// Completar el hábito
	return uc.registroRepo.CompletarHabito(idHabito, idUser, habito.GetPuntos())
}