// mascotas_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type MascotasRouter struct {
	engine *gin.Engine
}

func NewMascotasRouter(engine *gin.Engine) *MascotasRouter {
	return &MascotasRouter{
		engine: engine,
	}
}

func (router *MascotasRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByUserController,
	getByTipoMascotaController,
	getBySkinController,
	getByNivelController,
	updateExperienciaController,
	getMascotaCompletaController := InitMascotasDependencies()

	// Grupo de rutas para mascotas
	mascotasGroup := router.engine.Group("/mascotas")
	{
		// Rutas CRUD básicas
		mascotasGroup.POST("/", createController.Run)
		mascotasGroup.PUT("/:id", updateController.Run)
		mascotasGroup.DELETE("/:id", deleteController.Run)
		mascotasGroup.GET("/:id", getByIdController.Run)
		mascotasGroup.GET("/", getAllController.Run)
		
		// Rutas adicionales
		mascotasGroup.GET("/usuario/:id_user", getByUserController.Run)
		mascotasGroup.GET("/tipo-mascota/:id_tipo_mascota", getByTipoMascotaController.Run)
		mascotasGroup.GET("/skin/:id_skin", getBySkinController.Run)
		mascotasGroup.GET("/nivel/:id_nivel", getByNivelController.Run)
		
		// Rutas especiales
		mascotasGroup.PATCH("/:id/experiencia", updateExperienciaController.Run)
		mascotasGroup.GET("/:id/completa", getMascotaCompletaController.Run)
	}
}