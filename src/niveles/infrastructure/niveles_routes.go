// niveles_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type NivelesRouter struct {
	engine *gin.Engine
}

func NewNivelesRouter(engine *gin.Engine) *NivelesRouter {
	return &NivelesRouter{
		engine: engine,
	}
}

func (router *NivelesRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByNivelController,
	getByExpRequeridaController,
	getNextLevelController := InitNivelesDependencies()

	// Grupo de rutas para niveles
	nivelesGroup := router.engine.Group("/niveles")
	{
		// Rutas CRUD básicas
		nivelesGroup.POST("/", createController.Run)
		nivelesGroup.PUT("/:id", updateController.Run)
		nivelesGroup.DELETE("/:id", deleteController.Run)
		nivelesGroup.GET("/:id", getByIdController.Run)
		nivelesGroup.GET("/", getAllController.Run)
		
		// Rutas adicionales
		nivelesGroup.GET("/nivel/:nivel", getByNivelController.Run)
		nivelesGroup.GET("/exp/:exp", getByExpRequeridaController.Run)
		nivelesGroup.GET("/next-level", getNextLevelController.Run)
	}
}