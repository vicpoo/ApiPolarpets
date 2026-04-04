// habito_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type HabitoRouter struct {
	engine *gin.Engine
}

func NewHabitoRouter(engine *gin.Engine) *HabitoRouter {
	return &HabitoRouter{
		engine: engine,
	}
}

func (router *HabitoRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByUserController,
	getByTituloController,
	getByUserAndTituloController,
	getTotalPuntosController := InitHabitoDependencies()

	// Grupo de rutas para hábitos
	habitoGroup := router.engine.Group("/habitos")
	{
		// Rutas CRUD básicas
		habitoGroup.POST("/", createController.Run)
		habitoGroup.PUT("/:id", updateController.Run)
		habitoGroup.DELETE("/:id", deleteController.Run)
		habitoGroup.GET("/:id", getByIdController.Run)
		habitoGroup.GET("/", getAllController.Run)
		
		// Rutas adicionales
		habitoGroup.GET("/usuario/:id_user", getByUserController.Run)
		habitoGroup.GET("/usuario/:id_user/puntos", getTotalPuntosController.Run)
		
		// Rutas de búsqueda
		habitoGroup.GET("/search/titulo", getByTituloController.Run)
		habitoGroup.GET("/search", getByUserAndTituloController.Run)
	}
}