// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

func InitPagosDependencies() (
	*CreatePagoController,
	*UpdatePagoController,
	*DeletePagoController,
	*GetPagoByIdController,
	*GetAllPagosController,
	*GetPagosByUserController,
	*GetPagosByEstadoController,
	*GetPagosByMetodoPagoController,
	*GetPagosByFechaRangeController,
	*GetPagoByReferenciaExternaController,
	*GetTotalPagadoByUserController,
	*GetPagosCompletadosByUserController,
	*UpdatePagoEstadoController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLPagosRepository()

	// Casos de uso
	createUseCase := application.NewCreatePagoUseCase(repo)
	updateUseCase := application.NewUpdatePagoUseCase(repo)
	deleteUseCase := application.NewDeletePagoUseCase(repo)
	getByIdUseCase := application.NewGetPagoByIdUseCase(repo)
	getAllUseCase := application.NewGetAllPagosUseCase(repo)
	getByUserUseCase := application.NewGetPagosByUserUseCase(repo)
	getByEstadoUseCase := application.NewGetPagosByEstadoUseCase(repo)
	getByMetodoPagoUseCase := application.NewGetPagosByMetodoPagoUseCase(repo)
	getByFechaRangeUseCase := application.NewGetPagosByFechaRangeUseCase(repo)
	getByReferenciaExternaUseCase := application.NewGetPagoByReferenciaExternaUseCase(repo)
	getTotalPagadoUseCase := application.NewGetTotalPagadoByUserUseCase(repo)
	getPagosCompletadosUseCase := application.NewGetPagosCompletadosByUserUseCase(repo)
	updateEstadoUseCase := application.NewUpdatePagoEstadoUseCase(repo)

	// Controladores
	createController := NewCreatePagoController(createUseCase)
	updateController := NewUpdatePagoController(updateUseCase)
	deleteController := NewDeletePagoController(deleteUseCase)
	getByIdController := NewGetPagoByIdController(getByIdUseCase)
	getAllController := NewGetAllPagosController(getAllUseCase)
	getByUserController := NewGetPagosByUserController(getByUserUseCase)
	getByEstadoController := NewGetPagosByEstadoController(getByEstadoUseCase)
	getByMetodoPagoController := NewGetPagosByMetodoPagoController(getByMetodoPagoUseCase)
	getByFechaRangeController := NewGetPagosByFechaRangeController(getByFechaRangeUseCase)
	getByReferenciaExternaController := NewGetPagoByReferenciaExternaController(getByReferenciaExternaUseCase)
	getTotalPagadoController := NewGetTotalPagadoByUserController(getTotalPagadoUseCase)
	getPagosCompletadosController := NewGetPagosCompletadosByUserController(getPagosCompletadosUseCase)
	updateEstadoController := NewUpdatePagoEstadoController(updateEstadoUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByUserController,
		getByEstadoController,
		getByMetodoPagoController,
		getByFechaRangeController,
		getByReferenciaExternaController,
		getTotalPagadoController,
		getPagosCompletadosController,
		updateEstadoController
}