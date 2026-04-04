// pagos_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type PagosRouter struct {
	engine *gin.Engine
}

func NewPagosRouter(engine *gin.Engine) *PagosRouter {
	return &PagosRouter{
		engine: engine,
	}
}

func (router *PagosRouter) Run() {
	// Inicializar dependencias
	createController,
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
	updateEstadoController := InitPagosDependencies()

	// Grupo de rutas para pagos
	pagosGroup := router.engine.Group("/pagos")
	{
		// Rutas CRUD básicas
		pagosGroup.POST("/", createController.Run)
		pagosGroup.PUT("/:id", updateController.Run)
		pagosGroup.DELETE("/:id", deleteController.Run)
		pagosGroup.GET("/:id", getByIdController.Run)
		pagosGroup.GET("/", getAllController.Run)

		// Rutas por usuario
		pagosGroup.GET("/usuario/:id_usuario", getByUserController.Run)
		pagosGroup.GET("/usuario/:id_usuario/total", getTotalPagadoController.Run)
		pagosGroup.GET("/usuario/:id_usuario/completados", getPagosCompletadosController.Run)

		// Rutas de búsqueda y filtros
		pagosGroup.GET("/search/estado", getByEstadoController.Run)
		pagosGroup.GET("/search/metodo", getByMetodoPagoController.Run)
		pagosGroup.GET("/search/fecha", getByFechaRangeController.Run)
		pagosGroup.GET("/search/referencia", getByReferenciaExternaController.Run)

		// Rutas de acciones
		pagosGroup.PATCH("/:id/estado", updateEstadoController.Run)
	}
}