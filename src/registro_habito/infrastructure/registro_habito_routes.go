// registro_habito_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RegistroHabitoRouter struct {
	engine *gin.Engine
}

func NewRegistroHabitoRouter(engine *gin.Engine) *RegistroHabitoRouter {
	return &RegistroHabitoRouter{
		engine: engine,
	}
}

func (router *RegistroHabitoRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByHabitoController,
	getByHabitoAndFechaController,
	getByFechaRangeController,
	getByUserController,
	getTotalPuntosByHabitoController,
	getTotalPuntosByUserRegistroController,
	getRegistroCompletoController,
	getHabitosConEstadoController,
	completarHabitoController,
	existsRegistroHoyController := InitRegistroHabitoDependencies()

	// Grupo de rutas para registro de hábitos
	registroGroup := router.engine.Group("/registro-habito")
	{
		// Rutas CRUD básicas
		registroGroup.POST("/", createController.Run)
		registroGroup.PUT("/:id", updateController.Run)
		registroGroup.DELETE("/:id", deleteController.Run)
		registroGroup.GET("/:id", getByIdController.Run)
		registroGroup.GET("/", getAllController.Run)

		// Rutas adicionales
		registroGroup.GET("/habito/:id_habito", getByHabitoController.Run)
		registroGroup.GET("/usuario/:id_user", getByUserController.Run)
		registroGroup.GET("/usuario/:id_user/puntos", getTotalPuntosByUserRegistroController.Run)
		registroGroup.GET("/habito/:id_habito/puntos", getTotalPuntosByHabitoController.Run)
		registroGroup.GET("/:id/completo", getRegistroCompletoController.Run)
		
		// Rutas de búsqueda
		registroGroup.GET("/search", getByHabitoAndFechaController.Run)
		registroGroup.GET("/rango", getByFechaRangeController.Run)

		// Rutas para el FRONTEND (las más importantes)
		registroGroup.GET("/estado", getHabitosConEstadoController.Run)
		registroGroup.POST("/completar", completarHabitoController.Run)
		registroGroup.GET("/exists-hoy", existsRegistroHoyController.Run)
	}
}