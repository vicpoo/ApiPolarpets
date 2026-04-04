// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

func InitInventarioDependencies() (
	*CreateInventarioController,
	*UpdateInventarioController,
	*DeleteInventarioController,
	*DeleteInventarioByUserAndProductController,
	*GetInventarioByIdController,
	*GetAllInventarioController,
	*GetInventarioByUserController,
	*GetInventarioByProductoController,
	*GetInventarioByUserAndProductController,
	*ExistsInInventoryController,
	*GetInventarioByUserWithDetailsController,
	*GetInventarioByUserAndProductWithDetailsController,
	*GetCantidadProductosByUserController,
	*GetProductosByTipoInInventoryController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLInventarioUsuarioRepository()

	// Casos de uso
	createUseCase := application.NewCreateInventarioUseCase(repo)
	updateUseCase := application.NewUpdateInventarioUseCase(repo)
	deleteUseCase := application.NewDeleteInventarioUseCase(repo)
	deleteByUserAndProductUseCase := application.NewDeleteInventarioByUserAndProductUseCase(repo)
	getByIdUseCase := application.NewGetInventarioByIdUseCase(repo)
	getAllUseCase := application.NewGetAllInventarioUseCase(repo)
	getByUserUseCase := application.NewGetInventarioByUserUseCase(repo)
	getByProductoUseCase := application.NewGetInventarioByProductoUseCase(repo)
	getByUserAndProductUseCase := application.NewGetInventarioByUserAndProductUseCase(repo)
	existsInInventoryUseCase := application.NewExistsInInventoryUseCase(repo)
	getWithDetailsUseCase := application.NewGetInventarioByUserWithDetailsUseCase(repo)
	getWithDetailsByUserAndProductUseCase := application.NewGetInventarioByUserAndProductWithDetailsUseCase(repo)
	getCantidadUseCase := application.NewGetCantidadProductosByUserUseCase(repo)
	getByTipoUseCase := application.NewGetProductosByTipoInInventoryUseCase(repo)

	// Controladores
	createController := NewCreateInventarioController(createUseCase)
	updateController := NewUpdateInventarioController(updateUseCase)
	deleteController := NewDeleteInventarioController(deleteUseCase)
	deleteByUserAndProductController := NewDeleteInventarioByUserAndProductController(deleteByUserAndProductUseCase)
	getByIdController := NewGetInventarioByIdController(getByIdUseCase)
	getAllController := NewGetAllInventarioController(getAllUseCase)
	getByUserController := NewGetInventarioByUserController(getByUserUseCase)
	getByProductoController := NewGetInventarioByProductoController(getByProductoUseCase)
	getByUserAndProductController := NewGetInventarioByUserAndProductController(getByUserAndProductUseCase)
	existsInInventoryController := NewExistsInInventoryController(existsInInventoryUseCase)
	getWithDetailsController := NewGetInventarioByUserWithDetailsController(getWithDetailsUseCase)
	getWithDetailsByUserAndProductController := NewGetInventarioByUserAndProductWithDetailsController(getWithDetailsByUserAndProductUseCase)
	getCantidadController := NewGetCantidadProductosByUserController(getCantidadUseCase)
	getByTipoController := NewGetProductosByTipoInInventoryController(getByTipoUseCase)

	return createController,
		updateController,
		deleteController,
		deleteByUserAndProductController,
		getByIdController,
		getAllController,
		getByUserController,
		getByProductoController,
		getByUserAndProductController,
		existsInInventoryController,
		getWithDetailsController,
		getWithDetailsByUserAndProductController,
		getCantidadController,
		getByTipoController
}