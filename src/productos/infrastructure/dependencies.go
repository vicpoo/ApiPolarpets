// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

func InitProductosDependencies() (
	*CreateProductoController,
	*UpdateProductoController,
	*DeleteProductoController,
	*GetProductoByIdController,
	*GetAllProductosController,
	*GetProductosByTipoController,
	*GetProductosByPrecioRangeController,
	*GetProductosBySkinController,
	*GetProductosByTipoMascotaController,
	*GetProductoByNombreController,
	*GetProductoConDetallesController,
	*GetAllProductosConDetallesController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLProductosRepository()

	// Casos de uso
	createUseCase := application.NewCreateProductoUseCase(repo)
	updateUseCase := application.NewUpdateProductoUseCase(repo)
	deleteUseCase := application.NewDeleteProductoUseCase(repo)
	getByIdUseCase := application.NewGetProductoByIdUseCase(repo)
	getAllUseCase := application.NewGetAllProductosUseCase(repo)
	getByTipoUseCase := application.NewGetProductosByTipoUseCase(repo)
	getByPrecioRangeUseCase := application.NewGetProductosByPrecioRangeUseCase(repo)
	getBySkinUseCase := application.NewGetProductosBySkinUseCase(repo)
	getByTipoMascotaUseCase := application.NewGetProductosByTipoMascotaUseCase(repo)
	getByNombreUseCase := application.NewGetProductoByNombreUseCase(repo)
	getProductoConDetallesUseCase := application.NewGetProductoConDetallesUseCase(repo)
	getAllProductosConDetallesUseCase := application.NewGetAllProductosConDetallesUseCase(repo)

	// Controladores
	createController := NewCreateProductoController(createUseCase)
	updateController := NewUpdateProductoController(updateUseCase)
	deleteController := NewDeleteProductoController(deleteUseCase)
	getByIdController := NewGetProductoByIdController(getByIdUseCase)
	getAllController := NewGetAllProductosController(getAllUseCase)
	getByTipoController := NewGetProductosByTipoController(getByTipoUseCase)
	getByPrecioRangeController := NewGetProductosByPrecioRangeController(getByPrecioRangeUseCase)
	getBySkinController := NewGetProductosBySkinController(getBySkinUseCase)
	getByTipoMascotaController := NewGetProductosByTipoMascotaController(getByTipoMascotaUseCase)
	getByNombreController := NewGetProductoByNombreController(getByNombreUseCase)
	getProductoConDetallesController := NewGetProductoConDetallesController(getProductoConDetallesUseCase)
	getAllProductosConDetallesController := NewGetAllProductosConDetallesController(getAllProductosConDetallesUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByTipoController,
		getByPrecioRangeController,
		getBySkinController,
		getByTipoMascotaController,
		getByNombreController,
		getProductoConDetallesController,
		getAllProductosConDetallesController
}