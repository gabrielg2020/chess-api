package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielg2020/chess-api/api/handler/fen_handler"
	"github.com/gabrielg2020/chess-api/api/handler/move_handler"
	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gabrielg2020/chess-api/api/service/move_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for the root ("/") endpoint
func Test_Root_Endpoint(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	engine := setUpEngine()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Chess API!",
		})
	})

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err, "Expected Not to fail when generating mock request")
	rr := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `{"message":"Welcome to the Chess API!"}`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}

// Test for the /validate/fen endpoint
func Test_ValidateFEN_Endpoint(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	engine := setUpEngine()

	fenService := FENService.NewFENService()
	fenHandler := FENHandler.NewFENHandler(fenService)

	validateGroup := engine.Group("/validate")
	{
		validateGroup.GET("/fen", fenHandler.ValidateFEN)
	}

	req, err := http.NewRequest("GET", "/validate/fen?fen=rnbqkbnr%2Fpppppppp%2F8%2F8%2F8%2F8%2FPPPPPPPP%2FRNBQKBNR%20w%20KQkq%20-%200%201", nil)
	assert.NoError(t, err, "Expected Not to fail when generating mock request")
	rr := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `{"valid":true}`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}

func Test_GetMove_Endpoint(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	engine := setUpEngine()

	fenService := FENService.NewFENService()
	moveService := MoveService.NewMoveService()
	moveHandler := MoveHandler.NewMoveHandler(fenService, moveService)

	engine.GET("/move", moveHandler.FindBestMove)

	req, err := http.NewRequest("GET", "/move?fen=rnbqkbnr%2Fpppppppp%2F8%2F8%2F8%2F8%2FPPPPPPPP%2FRNBQKBNR%20w%20KQkq%20-%200%201", nil)
	assert.NoError(t, err, "Expected Not to fail when generating mock request")
	rr := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(rr, req)
	// Assert
	// TODO Complete test when move_service.FindBestMove is completed.
	//assert.Equal(t, http.StatusOK, rr.Code)
	// expectedBody := `{"move":"a2a4"}`
	// assert.JSONEq(t, expectedBody, rr.Body.String())

}
