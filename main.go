package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bamsy/go-echo-templ/handlers"
	customMiddleware "github.com/bamsy/go-echo-templ/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	gzip := middleware.GzipWithConfig(middleware.GzipConfig{Level: 5})

	landingHandler := handlers.LandingHandler{}

	// Create a group for the /static route and apply the gzipStatic middleware to it
	staticGroup := e.Group("/static")
	staticGroup.Use(customMiddleware.GzipStatic)
	staticGroup.Static("/", "static")

	e.GET("/", landingHandler.LandingPage, gzip)

	// check for dev mode
	if os.Getenv("DEV_MODE") == "true" {
		devMode(e)
	}

	startServer(e, ":8080")
}

// startServer starts the server on the specified port and gracefully shuts it down when an interrupt signal is received
func startServer(e *echo.Echo, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func devMode(e *echo.Echo) {
	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				// Read
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if err != nil {
					c.Logger().Error(err)
					return
				}

				// Write
				err = websocket.Message.Send(ws, "Hello, Client!")
				if err != nil {
					c.Logger().Error(err)
					return
				}
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
