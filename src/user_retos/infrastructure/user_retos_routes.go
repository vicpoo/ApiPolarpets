// user_retos_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UserRetosRouter struct {
	engine *gin.Engine
}

func NewUserRetosRouter(engine *gin.Engine) *UserRetosRouter {
	return &UserRetosRouter{
		engine: engine,
	}
}

func (router *UserRetosRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByUserController,
	getByRetoController,
	getByUserAndRetoController,
	getCompletedRetosController,
	getPendingRetosController,
	completeRetoController,
	getUserRetosConDetallesController := InitUserRetosDependencies()

	// Grupo de rutas para user_retos
	userRetosGroup := router.engine.Group("/user-retos")
	{
		// Rutas CRUD básicas
		userRetosGroup.POST("/", createController.Run)
		userRetosGroup.PUT("/:id", updateController.Run)
		userRetosGroup.DELETE("/:id", deleteController.Run)
		userRetosGroup.GET("/:id", getByIdController.Run)
		userRetosGroup.GET("/", getAllController.Run)

		// Rutas por usuario
		userRetosGroup.GET("/usuario/:id_usuario", getByUserController.Run)
		userRetosGroup.GET("/usuario/:id_usuario/completados", getCompletedRetosController.Run)
		userRetosGroup.GET("/usuario/:id_usuario/pendientes", getPendingRetosController.Run)
		userRetosGroup.GET("/usuario/:id_usuario/detalles", getUserRetosConDetallesController.Run)

		// Rutas por reto
		userRetosGroup.GET("/reto/:id_reto", getByRetoController.Run)

		// Rutas de búsqueda y acciones
		userRetosGroup.GET("/search", getByUserAndRetoController.Run)
		userRetosGroup.POST("/completar", completeRetoController.Run)
	}
}