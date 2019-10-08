package main

import (
	"api_graphql/handlers"
	"github.com/labstack/echo"
)

func main() {
	api := echo.New()
	api.GET("/", handlers.IndexGet)
	api.POST("/", handlers.IndexPost)
	api.PUT("/", handlers.IndexPut)
	api.DELETE("/", handlers.IndexDelete)

	api.GET("/user", handlers.UserGet)
	api.POST("/user", handlers.UserPost)
	api.PUT("/user", handlers.UserPut)
	api.DELETE("/user", handlers.UserDelete)

	api.Logger.Error(api.Start(":1000"))
}
