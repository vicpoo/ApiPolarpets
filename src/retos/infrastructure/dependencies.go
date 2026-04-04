// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
)

func InitRetosDependencies() (
	*CreateRetoController,
	*UpdateRetoController,
	*DeleteRetoController,
	*GetRetoByIdController,
	*GetAllRetosController,
	*GetRetoByTituloController,
	*GetRetosByPuntosRangeController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLRetosRepository()

	// Casos de uso
	createUseCase := application.NewCreateRetoUseCase(repo)
	updateUseCase := application.NewUpdateRetoUseCase(repo)
	deleteUseCase := application.NewDeleteRetoUseCase(repo)
	getByIdUseCase := application.NewGetRetoByIdUseCase(repo)
	getAllUseCase := application.NewGetAllRetosUseCase(repo)
	getByTituloUseCase := application.NewGetRetoByTituloUseCase(repo)
	getByPuntosRangeUseCase := application.NewGetRetosByPuntosRangeUseCase(repo)

	// Controladores
	createController := NewCreateRetoController(createUseCase)
	updateController := NewUpdateRetoController(updateUseCase)
	deleteController := NewDeleteRetoController(deleteUseCase)
	getByIdController := NewGetRetoByIdController(getByIdUseCase)
	getAllController := NewGetAllRetosController(getAllUseCase)
	getByTituloController := NewGetRetoByTituloController(getByTituloUseCase)
	getByPuntosRangeController := NewGetRetosByPuntosRangeController(getByPuntosRangeUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByTituloController,
		getByPuntosRangeController
}