package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/core"

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

	
	port := ":8000"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	if err := engine.Run(port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("No se puede iniciar el servidor : %v", err)
	}
}