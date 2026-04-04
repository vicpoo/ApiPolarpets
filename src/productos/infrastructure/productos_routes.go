// productos_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type ProductosRouter struct {
	engine *gin.Engine
}

func NewProductosRouter(engine *gin.Engine) *ProductosRouter {
	return &ProductosRouter{
		engine: engine,
	}
}

func (router *ProductosRouter) Run() {
	// Inicializar dependencias
	createController,
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
	getAllProductosConDetallesController := InitProductosDependencies()

	// Grupo de rutas para productos
	productosGroup := router.engine.Group("/productos")
	{
		// Rutas CRUD básicas
		productosGroup.POST("/", createController.Run)
		productosGroup.PUT("/:id", updateController.Run)
		productosGroup.DELETE("/:id", deleteController.Run)
		productosGroup.GET("/:id", getByIdController.Run)
		productosGroup.GET("/", getAllController.Run)

		// Rutas de búsqueda y filtros
		productosGroup.GET("/search/tipo", getByTipoController.Run)
		productosGroup.GET("/search/precio", getByPrecioRangeController.Run)
		productosGroup.GET("/search/nombre", getByNombreController.Run)

		// Rutas por relaciones
		productosGroup.GET("/skin/:id_skin", getBySkinController.Run)
		productosGroup.GET("/tipo-mascota/:id_tipo_mascota", getByTipoMascotaController.Run)

		// Rutas con detalles (joins)
		productosGroup.GET("/detalles", getAllProductosConDetallesController.Run)
		productosGroup.GET("/:id/detalles", getProductoConDetallesController.Run)
	}
}