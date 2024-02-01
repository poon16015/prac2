package main

import (
	"github.com/gin-gonic/gin"
	"github.com/poon16015/prac2/controller"
	"github.com/poon16015/prac2/entity"
)

func main() {
	entity.ConnectDB()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Auth Routes
	router := r.Group("")
	{
		// router.Use(middlewares.Authorizes())
		// {
		// User Routes
		router.GET("/users", controller.ListUsers)
		router.GET("/user/:id", controller.GetUser)
		router.POST("/users", controller.CreateUser)
		router.PATCH("/users", controller.UpdateUser)
		router.DELETE("/users/:id", controller.DeleteUser)
		
	}

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
