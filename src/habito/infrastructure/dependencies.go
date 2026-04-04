// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

func InitHabitoDependencies() (
	*CreateHabitoController,
	*UpdateHabitoController,
	*DeleteHabitoController,
	*GetHabitoByIdController,
	*GetAllHabitosController,
	*GetHabitosByUserController,
	*GetHabitoByTituloController,
	*GetHabitoByUserAndTituloController,
	*GetTotalPuntosByUserController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLHabitoRepository()

	// Casos de uso
	createUseCase := application.NewCreateHabitoUseCase(repo)
	updateUseCase := application.NewUpdateHabitoUseCase(repo)
	deleteUseCase := application.NewDeleteHabitoUseCase(repo)
	getByIdUseCase := application.NewGetHabitoByIdUseCase(repo)
	getAllUseCase := application.NewGetAllHabitosUseCase(repo)
	getByUserUseCase := application.NewGetHabitosByUserUseCase(repo)
	getByTituloUseCase := application.NewGetHabitoByTituloUseCase(repo)
	getByUserAndTituloUseCase := application.NewGetHabitoByUserAndTituloUseCase(repo)
	getTotalPuntosUseCase := application.NewGetTotalPuntosByUserUseCase(repo)

	// Controladores
	createController := NewCreateHabitoController(createUseCase)
	updateController := NewUpdateHabitoController(updateUseCase)
	deleteController := NewDeleteHabitoController(deleteUseCase)
	getByIdController := NewGetHabitoByIdController(getByIdUseCase)
	getAllController := NewGetAllHabitosController(getAllUseCase)
	getByUserController := NewGetHabitosByUserController(getByUserUseCase)
	getByTituloController := NewGetHabitoByTituloController(getByTituloUseCase)
	getByUserAndTituloController := NewGetHabitoByUserAndTituloController(getByUserAndTituloUseCase)
	getTotalPuntosController := NewGetTotalPuntosByUserController(getTotalPuntosUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByUserController,
		getByTituloController,
		getByUserAndTituloController,
		getTotalPuntosController
}