// inventario_usuario_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type InventarioRouter struct {
	engine *gin.Engine
}

func NewInventarioRouter(engine *gin.Engine) *InventarioRouter {
	return &InventarioRouter{
		engine: engine,
	}
}

func (router *InventarioRouter) Run() {
	// Inicializar dependencias
	createController,
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
	getByTipoController := InitInventarioDependencies()

	// Grupo de rutas para inventario de usuario
	inventarioGroup := router.engine.Group("/inventario")
	{
		// Rutas CRUD básicas
		inventarioGroup.POST("/", createController.Run)
		inventarioGroup.PUT("/:id", updateController.Run)
		inventarioGroup.DELETE("/:id", deleteController.Run)
		inventarioGroup.GET("/:id", getByIdController.Run)
		inventarioGroup.GET("/", getAllController.Run)

		// Rutas por usuario
		inventarioGroup.GET("/usuario/:id_usuario", getByUserController.Run)
		inventarioGroup.GET("/usuario/:id_usuario/detalles", getWithDetailsController.Run)
		inventarioGroup.GET("/usuario/:id_usuario/cantidad", getCantidadController.Run)
		inventarioGroup.GET("/usuario/:id_usuario/tipo", getByTipoController.Run)

		// Rutas por producto
		inventarioGroup.GET("/producto/:id_producto", getByProductoController.Run)

		// Rutas de eliminación específica
		inventarioGroup.DELETE("/usuario/:id_usuario/producto/:id_producto", deleteByUserAndProductController.Run)

		// Rutas de búsqueda y verificación
		inventarioGroup.GET("/search", getByUserAndProductController.Run)
		inventarioGroup.GET("/exists", existsInInventoryController.Run)

		// Ruta con detalles específicos
		inventarioGroup.GET("/detalles", getWithDetailsByUserAndProductController.Run)
	}
}