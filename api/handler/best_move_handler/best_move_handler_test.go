package BestMoveHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielg2020/chess-api/api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_BestMoveHandler_FindBestMove(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name                 string
		fen                  string
		setupFENService      func(m *mocks.MockFENService)
		setupBestMoveService func(m *mocks.MockBestMoveService)
		expectedStatusCode   int
		expectedResponse     string
	}{
		{
			name: "Test Case 1",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			setupFENService: func(m *mocks.MockFENService) {
				m.On("Validate", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(nil)
				mockChessboard := new(mocks.MockChessboardEntity)
				m.On("Parse", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(mockChessboard, nil)
			},
			setupBestMoveService: func(m *mocks.MockBestMoveService) {
				m.On("FindBestMove", mock.Anything).Return("a2a4", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"bestMove":"a2a4"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockFENService := new(mocks.MockFENService)
			mockBestMoveService := new(mocks.MockBestMoveService)
			handler := NewBestMoveHandler(mockFENService, mockBestMoveService)
			tc.setupFENService(mockFENService)
			tc.setupBestMoveService(mockBestMoveService)

			engine := gin.Default()
			engine.GET("/best_move", handler.FindBestMove)
			req, err := http.NewRequest(http.MethodGet, "/best_move?fen="+tc.fen, nil)
			assert.NoError(t, err, "Expected Not to fail when generating mock request")

			rr := httptest.NewRecorder()

			// Act
			engine.ServeHTTP(rr, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode, "Expected mock status code and test status code to equate")
			assert.JSONEq(t, tc.expectedResponse, rr.Body.String(), "Expected mock response and test response to equate")
			mockFENService.AssertNumberOfCalls(t, "Validate", 1)
			mockFENService.AssertNumberOfCalls(t, "Parse", 1)
			mockBestMoveService.AssertNumberOfCalls(t, "FindBestMove", 1)
		})
	}
}

