package main

import (
	"log"

	"anima.onl/render"
	"anima.onl/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.HTMLRender = render.Create()

	routes.RegisterRoutes(e)

	if err := e.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
