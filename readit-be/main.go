package main

import (
	"readit-be/controllers"
	"readit-be/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	database.ConnectDatabase()
	r.Static("/assets", "./ui/assets")
	r.GET("/", func(c *gin.Context) {
		c.File("./ui/index.html")
	})
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/index.html")
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)

	// >> These are protected routes
	auth: r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware()){
		r.GET("/posts", controllers.GetPosts)
		r.POST("/posts", controllers.CreatePost)
	}
	
	// r.Static("/assets", "./ui/dist/assets")
	// r.GET("/", func(c *gin.Context) {
	// 	c.File("./ui/dist/index.html")
	// })
	// r.NoRoute(func(c *gin.Context) {
	// 	c.File("./ui/dist/index.html")
	// })

	r.Run(":8080")
}
