package routes

import (
	"gellyzxc-template-golang-gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	//r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)
}
