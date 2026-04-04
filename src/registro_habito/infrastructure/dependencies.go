// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
	habitoInfra "github.com/vicpoo/ApiPolarpets/src/habito/infrastructure"
)

func InitRegistroHabitoDependencies() (
	*CreateRegistroHabitoController,
	*UpdateRegistroHabitoController,
	*DeleteRegistroHabitoController,
	*GetRegistroHabitoByIdController,
	*GetAllRegistroHabitosController,
	*GetRegistroHabitoByHabitoController,
	*GetRegistroHabitoByHabitoAndFechaController,
	*GetRegistroHabitoByFechaRangeController,
	*GetRegistroHabitoByUserController,
	*GetTotalPuntosByHabitoController,
	*GetTotalPuntosByUserRegistroController,
	*GetRegistroCompletoController,
	*GetHabitosConEstadoController,
	*CompletarHabitoController,
	*ExistsRegistroHoyController,
) {
	// Repositorios
	registroRepo := NewMySQLRegistroHabitoRepository()
	habitoRepo := habitoInfra.NewMySQLHabitoRepository()

	// Casos de uso - CRUD básico
	createUseCase := application.NewCreateRegistroHabitoUseCase(registroRepo)
	updateUseCase := application.NewUpdateRegistroHabitoUseCase(registroRepo)
	deleteUseCase := application.NewDeleteRegistroHabitoUseCase(registroRepo)
	getByIdUseCase := application.NewGetRegistroHabitoByIdUseCase(registroRepo)
	getAllUseCase := application.NewGetAllRegistroHabitosUseCase(registroRepo)

	// Casos de uso - Métodos adicionales
	getByHabitoUseCase := application.NewGetRegistroHabitoByHabitoUseCase(registroRepo)
	getByHabitoAndFechaUseCase := application.NewGetRegistroHabitoByHabitoAndFechaUseCase(registroRepo)
	getByFechaRangeUseCase := application.NewGetRegistroHabitoByFechaRangeUseCase(registroRepo)
	getByUserUseCase := application.NewGetRegistroHabitoByUserUseCase(registroRepo)
	getTotalPuntosByHabitoUseCase := application.NewGetTotalPuntosByHabitoUseCase(registroRepo)
	getTotalPuntosByUserUseCase := application.NewGetTotalPuntosByUserRegistroUseCase(registroRepo)
	getRegistroCompletoUseCase := application.NewGetRegistroCompletoUseCase(registroRepo)

	// Casos de uso - Frontend
	getHabitosConEstadoUseCase := application.NewGetHabitosConEstadoUseCase(registroRepo)
	completarHabitoUseCase := application.NewCompletarHabitoUseCase(registroRepo, habitoRepo)
	existsRegistroHoyUseCase := application.NewExistsRegistroHoyUseCase(registroRepo)

	// Controladores - CRUD básico
	createController := NewCreateRegistroHabitoController(createUseCase)
	updateController := NewUpdateRegistroHabitoController(updateUseCase)
	deleteController := NewDeleteRegistroHabitoController(deleteUseCase)
	getByIdController := NewGetRegistroHabitoByIdController(getByIdUseCase)
	getAllController := NewGetAllRegistroHabitosController(getAllUseCase)

	// Controladores - Métodos adicionales
	getByHabitoController := NewGetRegistroHabitoByHabitoController(getByHabitoUseCase)
	getByHabitoAndFechaController := NewGetRegistroHabitoByHabitoAndFechaController(getByHabitoAndFechaUseCase)
	getByFechaRangeController := NewGetRegistroHabitoByFechaRangeController(getByFechaRangeUseCase)
	getByUserController := NewGetRegistroHabitoByUserController(getByUserUseCase)
	getTotalPuntosByHabitoController := NewGetTotalPuntosByHabitoController(getTotalPuntosByHabitoUseCase)
	getTotalPuntosByUserRegistroController := NewGetTotalPuntosByUserRegistroController(getTotalPuntosByUserUseCase)
	getRegistroCompletoController := NewGetRegistroCompletoController(getRegistroCompletoUseCase)

	// Controladores - Frontend
	getHabitosConEstadoController := NewGetHabitosConEstadoController(getHabitosConEstadoUseCase)
	completarHabitoController := NewCompletarHabitoController(completarHabitoUseCase)
	existsRegistroHoyController := NewExistsRegistroHoyController(existsRegistroHoyUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByHabitoController,
		getByHabitoAndFechaController,
		getByFechaRangeController,
		getByUserController,
		getTotalPuntosByHabitoController,
		getTotalPuntosByUserRegistroController,
		getRegistroCompletoController,
		getHabitosConEstadoController,
		completarHabitoController,
		existsRegistroHoyController
}