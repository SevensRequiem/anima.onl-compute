package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {

	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{"title": "Home"})
	})

	apiGroup := e.Group("/api")
	{
		apiGroup.GET("/users", getUsers)
		apiGroup.POST("/users", createUser)
	}
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": []string{"Alice", "Bob"}})
}

func createUser(c *gin.Context) {
	c.Status(http.StatusCreated)
}
