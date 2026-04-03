// usuario_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UsuarioRouter struct {
	engine *gin.Engine
}

func NewUsuarioRouter(engine *gin.Engine) *UsuarioRouter {
	return &UsuarioRouter{
		engine: engine,
	}
}

func (router *UsuarioRouter) Run() {
	// Inicializar dependencias
	registerController,
	loginController,
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByEmailController,
	updateMascotaActivaController := InitUsuarioDependencies()

	// Grupo de rutas para usuarios
	usuarioGroup := router.engine.Group("/usuarios")
	{
		// Rutas de autenticación
		usuarioGroup.POST("/register", registerController.Run)
		usuarioGroup.POST("/login", loginController.Run)
		
		// Rutas CRUD
		usuarioGroup.POST("/", createController.Run)
		usuarioGroup.PUT("/:id", updateController.Run)
		usuarioGroup.DELETE("/:id", deleteController.Run)
		usuarioGroup.GET("/:id", getByIdController.Run)
		usuarioGroup.GET("/", getAllController.Run)
		
		// Rutas adicionales
		usuarioGroup.GET("/search/email", getByEmailController.Run)
		usuarioGroup.PATCH("/:id/mascota-activa", updateMascotaActivaController.Run)
	}
}