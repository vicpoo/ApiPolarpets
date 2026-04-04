// skins_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type SkinsRouter struct {
	engine *gin.Engine
}

func NewSkinsRouter(engine *gin.Engine) *SkinsRouter {
	return &SkinsRouter{
		engine: engine,
	}
}

func (router *SkinsRouter) Run() {
	// Inicializar dependencias
	createController,
	updateController,
	deleteController,
	getByIdController,
	getAllController,
	getByTipoMascotaController,
	getByNombreController,
	getByTipoMascotaAndNombreController := InitSkinsDependencies()

	// Grupo de rutas para skins
	skinsGroup := router.engine.Group("/skins")
	{
		// Rutas CRUD básicas
		skinsGroup.POST("/", createController.Run)
		skinsGroup.PUT("/:id", updateController.Run)
		skinsGroup.DELETE("/:id", deleteController.Run)
		skinsGroup.GET("/:id", getByIdController.Run)
		skinsGroup.GET("/", getAllController.Run)
		
		// Rutas adicionales
		skinsGroup.GET("/tipo-mascota/:id_tipo_mascota", getByTipoMascotaController.Run)
		skinsGroup.GET("/search/nombre", getByNombreController.Run)
		skinsGroup.GET("/search", getByTipoMascotaAndNombreController.Run)
	}
}