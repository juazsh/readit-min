package main

import (
	"readit-be/controllers"
	"readit-be/database"
	"readit-be/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // your Vite frontend
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/posts", controllers.GetPosts)
		auth.POST("/posts", controllers.CreatePost)
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
