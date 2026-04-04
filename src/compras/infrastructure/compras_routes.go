// compras_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type ComprasRouter struct {
	engine *gin.Engine
}

func NewComprasRouter(engine *gin.Engine) *ComprasRouter {
	return &ComprasRouter{
		engine: engine,
	}
}

func (router *ComprasRouter) Run() {
	// Inicializar dependencias
	createController,
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
	getRecientesController := InitComprasDependencies()

	// Grupo de rutas para compras
	comprasGroup := router.engine.Group("/compras")
	{
		// Rutas CRUD básicas
		comprasGroup.POST("/", createController.Run)
		comprasGroup.PUT("/:id", updateController.Run)
		comprasGroup.DELETE("/:id", deleteController.Run)
		comprasGroup.GET("/:id", getByIdController.Run)
		comprasGroup.GET("/", getAllController.Run)

		// Rutas por usuario
		comprasGroup.GET("/usuario/:id_usuario", getByUserController.Run)
		comprasGroup.GET("/usuario/:id_usuario/detalles", getWithDetailsController.Run)
		comprasGroup.GET("/usuario/:id_usuario/total", getTotalGastadoController.Run)
		comprasGroup.GET("/usuario/:id_usuario/recientes", getRecientesController.Run)

		// Rutas por producto y pago
		comprasGroup.GET("/producto/:id_producto", getByProductoController.Run)
		comprasGroup.GET("/pago/:id_pago", getByPagoController.Run)

		// Rutas de búsqueda
		comprasGroup.GET("/search/fecha", getByFechaRangeController.Run)

		// Ruta con detalles específicos
		comprasGroup.GET("/:id/detalles", getByIdWithDetailsController.Run)
	}
}