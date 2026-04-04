// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

func InitComprasDependencies() (
	*CreateCompraController,
	*UpdateCompraController,
	*DeleteCompraController,
	*GetCompraByIdController,
	*GetAllComprasController,
	*GetComprasByUserController,
	*GetComprasByProductoController,
	*GetCompraByPagoController,
	*GetComprasByFechaRangeController,
	*GetComprasByUserWithDetailsController,
	*GetCompraByIdWithDetailsController,
	*GetTotalGastadoByUserController,
	*GetComprasRecientesByUserController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLComprasRepository()

	// Casos de uso
	createUseCase := application.NewCreateCompraUseCase(repo)
	updateUseCase := application.NewUpdateCompraUseCase(repo)
	deleteUseCase := application.NewDeleteCompraUseCase(repo)
	getByIdUseCase := application.NewGetCompraByIdUseCase(repo)
	getAllUseCase := application.NewGetAllComprasUseCase(repo)
	getByUserUseCase := application.NewGetComprasByUserUseCase(repo)
	getByProductoUseCase := application.NewGetComprasByProductoUseCase(repo)
	getByPagoUseCase := application.NewGetCompraByPagoUseCase(repo)
	getByFechaRangeUseCase := application.NewGetComprasByFechaRangeUseCase(repo)
	getWithDetailsUseCase := application.NewGetComprasByUserWithDetailsUseCase(repo)
	getByIdWithDetailsUseCase := application.NewGetCompraByIdWithDetailsUseCase(repo)
	getTotalGastadoUseCase := application.NewGetTotalGastadoByUserUseCase(repo)
	getRecientesUseCase := application.NewGetComprasRecientesByUserUseCase(repo)

	// Controladores
	createController := NewCreateCompraController(createUseCase)
	updateController := NewUpdateCompraController(updateUseCase)
	deleteController := NewDeleteCompraController(deleteUseCase)
	getByIdController := NewGetCompraByIdController(getByIdUseCase)
	getAllController := NewGetAllComprasController(getAllUseCase)
	getByUserController := NewGetComprasByUserController(getByUserUseCase)
	getByProductoController := NewGetComprasByProductoController(getByProductoUseCase)
	getByPagoController := NewGetCompraByPagoController(getByPagoUseCase)
	getByFechaRangeController := NewGetComprasByFechaRangeController(getByFechaRangeUseCase)
	getWithDetailsController := NewGetComprasByUserWithDetailsController(getWithDetailsUseCase)
	getByIdWithDetailsController := NewGetCompraByIdWithDetailsController(getByIdWithDetailsUseCase)
	getTotalGastadoController := NewGetTotalGastadoByUserController(getTotalGastadoUseCase)
	getRecientesController := NewGetComprasRecientesByUserController(getRecientesUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByUserController,
		getByProductoController,
		getByPagoController,
		getByFechaRangeController,
		getWithDetailsController,
		getByIdWithDetailsController,
		getTotalGastadoController,
		getRecientesController
}