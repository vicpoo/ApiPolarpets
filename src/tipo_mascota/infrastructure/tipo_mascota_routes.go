// tipo_mascota_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type TipoMascotaRouter struct {
	engine *gin.Engine
}

func NewTipoMascotaRouter(engine *gin.Engine) *TipoMascotaRouter {
	return &TipoMascotaRouter{
		engine: engine,
	}
}

func (router *TipoMascotaRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByNombreController := InitTipoMascotaDependencies()

	// Grupo de rutas para tipos de mascota
	tipoMascotaGroup := router.engine.Group("/tipos-mascota")
	{
		// Rutas CRUD básicas
		tipoMascotaGroup.POST("/", createController.Run)
		tipoMascotaGroup.PUT("/:id", updateController.Run)
		tipoMascotaGroup.DELETE("/:id", deleteController.Run)
		tipoMascotaGroup.GET("/:id", getByIdController.Run)
		tipoMascotaGroup.GET("/", getAllController.Run)
		
		// Ruta adicional
		tipoMascotaGroup.GET("/search", getByNombreController.Run)
	}
}