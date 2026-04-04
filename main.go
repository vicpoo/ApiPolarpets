package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/core"


	userInfrastructure "github.com/vicpoo/ApiPolarpets/src/usuarios/infrastructure"
	rolInfrastructure "github.com/vicpoo/ApiPolarpets/src/roles/infrastructure"
	comprasInfrastructure "github.com/vicpoo/ApiPolarpets/src/compras/infrastructure"
	habitoInfrastructure "github.com/vicpoo/ApiPolarpets/src/habito/infrastructure"
	inventario_usuarioInfrastructure "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/infrastructure"
	mascotasInfrastructure "github.com/vicpoo/ApiPolarpets/src/mascotas/infrastructure"
	nivelesInfrastructure "github.com/vicpoo/ApiPolarpets/src/niveles/infrastructure"
	pagosInfrastructure "github.com/vicpoo/ApiPolarpets/src/pagos/infrastructure"
	productosInfrastructure "github.com/vicpoo/ApiPolarpets/src/productos/infrastructure"
	registro_habitoInfrastructure "github.com/vicpoo/ApiPolarpets/src/registro_habito/infrastructure"
	retosInfrastructure "github.com/vicpoo/ApiPolarpets/src/retos/infrastructure"
	skinsInfrastructure "github.com/vicpoo/ApiPolarpets/src/skins/infrastructure"
	tipo_mascotaInfrastructure "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/infrastructure"
	user_retosInfrastructure "github.com/vicpoo/ApiPolarpets/src/user_retos/infrastructure"


)

func main(){
	core.InitDB()

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))


	userRouter := userInfrastructure.NewUsuarioRouter(engine)
	userRouter.Run()

	rolRouter := rolInfrastructure.NewRolRouter(engine)
	rolRouter.Run()

	comprasRouter := comprasInfrastructure.NewComprasRouter(engine)
	comprasRouter.Run()

	habitoRouter := habitoInfrastructure.NewHabitoRouter(engine)
	habitoRouter.Run()

	inventario_usuarioRouter := inventario_usuarioInfrastructure.NewInventarioRouter(engine)
	inventario_usuarioRouter.Run()

	mascotasRouter := mascotasInfrastructure.NewMascotasRouter(engine)
	mascotasRouter.Run()

	nivelesRouter := nivelesInfrastructure.NewNivelesRouter(engine)
	nivelesRouter.Run()

	pagosRouter := pagosInfrastructure.NewPagosRouter(engine)
	pagosRouter.Run()

	productosRouter := productosInfrastructure.NewProductosRouter(engine)
	productosRouter.Run()

	registro_habitoRouter := registro_habitoInfrastructure.NewRegistroHabitoRouter(engine)
	registro_habitoRouter.Run()

	retosRouter := retosInfrastructure.NewRetosRouter(engine)
	retosRouter.Run()

	skinsRouter := skinsInfrastructure.NewSkinsRouter(engine)
	skinsRouter.Run()

	tipo_mascotaRouter := tipo_mascotaInfrastructure.NewTipoMascotaRouter(engine)
	tipo_mascotaRouter.Run()

	user_retosRouter := user_retosInfrastructure.NewUserRetosRouter(engine)
	user_retosRouter.Run()


	
	port := ":8000"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	if err := engine.Run(port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("No se puede iniciar el servidor : %v", err)
	}
}