package MoveHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielg2020/chess-api/api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_MoveHandler_FindBestMove(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name               string
		fen                string
		setupFENService    func(m *mocks.MockFENService)
		setupMoveService   func(m *mocks.MockMoveService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "Test Case 1",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			setupFENService: func(m *mocks.MockFENService) {
				m.On("Validate", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(nil)
				mockChessboard := new(mocks.MockChessboardEntity)
				m.On("Parse", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(mockChessboard, nil)
			},
			setupMoveService: func(m *mocks.MockMoveService) {
				m.On("FindBestMove", mock.Anything).Return("a2a4", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"move":"a2a4"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockFENService := new(mocks.MockFENService)
			mockMoveService := new(mocks.MockMoveService)
			handler := NewMoveHandler(mockFENService, mockMoveService)
			tc.setupFENService(mockFENService)
			tc.setupMoveService(mockMoveService)

			engine := gin.Default()
			engine.GET("/move", handler.FindBestMove)
			req, err := http.NewRequest(http.MethodGet, "/move?fen="+tc.fen, nil)
			assert.NoError(t, err, "Expected Not to fail when generating mock request")

			rr := httptest.NewRecorder()

			// Act
			engine.ServeHTTP(rr, req)

			// Assert
			// TODO [FindBestMove] Complete test when move_service.FindBestMove is completed.
			// assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode, "Expected mock status code and test status code to equate")
			// assert.JSONEq(t, tc.expectedResponse, rr.Body.String(), "Expected mock response and test response to equate")
			// mockFENService.AssertNumberOfCalls(t, "Validate", 1)
			// mockFENService.AssertNumberOfCalls(t, "Parse", 1)
			// mockMoveService.AssertNumberOfCalls(t, "FindBestMove", 1)
		})
	}
}
