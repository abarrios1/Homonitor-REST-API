package main

import (
	"github.com/abarrios1/HomonitorBackend/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	echoRouter := echo.New()

	// Middleware
	echoRouter.Use(middleware.Logger()) // Server log
	echoRouter.Use(middleware.Recover()) // Recover from panics anywhere in the chain

	// CORS -- Cross Domain Access Controls
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET,echo.HEAD,echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	echoRouter.GET("/", hello)
	echoRouter.PUT("/devices/create", controllers.CreateDevice)
	echoRouter.POST("/devices/information", controllers.DeviceInterfaceEntryStats)

	// Start server
	echoRouter.Logger.Fatal(echoRouter.Start(":9000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
