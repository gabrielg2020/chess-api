package main

import (
	"github.com/gabrielg2020/chess-api/api/handler/fen_handler"
	"github.com/gabrielg2020/chess-api/api/handler/move_handler"
	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gabrielg2020/chess-api/api/service/move_service"
	"github.com/gabrielg2020/chess-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

// FEAT Add a way to propagate error messages. Maybe add a logger like logrus?
func main() {
	engine := setUpEngine()

	// Initialise services
	fenService := FENService.NewFENService()
	moveService := MoveService.NewMoveService()

	// Initialise handlers
	fenHandler := FENHandler.NewFENHandler(fenService)
	moveHandler := MoveHandler.NewMoveHandler(fenService, moveService)

	// Set up endpoints
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Chess API!",
		})
	})

	validateGroup := engine.Group("/validate")
	{
		validateGroup.GET("/fen", fenHandler.ValidateFEN)
	}

	engine.GET("/move", moveHandler.FindBestMove)

	// Start engine
	if err := engine.Run(":8080"); err != nil {
		logger.Log.WithError(err).Fatal("Failed to start server")
	}
}

func setUpEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Log.WithError(err).Fatal("Failed to set trusted proxies")
	}

	return engine
}
