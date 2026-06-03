package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/meplos/tetris/internal/leaderboard"
)

func main() {
	server := echo.New()

	server.Use(middleware.RequestLogger())

	server.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "healthy",
		})
	})

	leaderboardHandler := leaderboard.NewHandler()

	leaderboardHandler.Setup(server.Group("/api"))

	server.StaticFS("/public", os.DirFS("./static"))

	if err := server.Start(":9090"); err != nil {
		log.Fatalf("error: failed to start server %s", err)
	}
}
