// retos_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RetosRouter struct {
	engine *gin.Engine
}

func NewRetosRouter(engine *gin.Engine) *RetosRouter {
	return &RetosRouter{
		engine: engine,
	}
}

func (router *RetosRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByTituloController,
	getByPuntosRangeController := InitRetosDependencies()

	// Grupo de rutas para retos
	retosGroup := router.engine.Group("/retos")
	{
		// Rutas CRUD básicas
		retosGroup.POST("/", createController.Run)
		retosGroup.PUT("/:id", updateController.Run)
		retosGroup.DELETE("/:id", deleteController.Run)
		retosGroup.GET("/:id", getByIdController.Run)
		retosGroup.GET("/", getAllController.Run)

		// Rutas de búsqueda
		retosGroup.GET("/search/titulo", getByTituloController.Run)
		retosGroup.GET("/search/rango", getByPuntosRangeController.Run)
	}
}