package FENHandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielg2020/chess-api/api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_FENHandler_ValidateFEN(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name               string
		fen                string
		setupFENService    func(m *mocks.MockFENService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "Valid FEN",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			setupFENService: func(m *mocks.MockFENService) {
				m.On("Validate", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"valid":true}`,
		},
		{
			name: "Invalid FEN [FEN string doesn't pass regex]",
			fen:  "fen",
			setupFENService: func(m *mocks.MockFENService) {
				m.On("Validate", "fen").Return(errors.New("FENService.Validate: string is not a FEN"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"valid":false, "errorMessage":"inputted fen is invalid", "errorCode":400}`,
		},
		{
			name: "Invalid FEN [FEN string is empty]",
			fen:  "",
			setupFENService: func(m *mocks.MockFENService) {
				m.On("Validate", "").Return(errors.New("FENService.Validate: FEN string empty"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"valid":false, "errorMessage":"inputted fen is invalid", "errorCode":400}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			mockFENService := new(mocks.MockFENService)
			handler := NewFENHandler(mockFENService)
			tc.setupFENService(mockFENService)

			engine := gin.Default()
			engine.GET("/validate/fen", handler.ValidateFEN)
			req, err := http.NewRequest(http.MethodGet, "/validate/fen?fen="+tc.fen, nil)
			assert.NoError(t, err, "Expected Not to fail when generating mock request")

			rr := httptest.NewRecorder()

			// Act
			engine.ServeHTTP(rr, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode, "Expected mock status code and test status code to equate")
			assert.JSONEq(t, tc.expectedResponse, rr.Body.String(), "Expected mock response and test response to equate")
			mockFENService.AssertCalled(t, "Validate", tc.fen)
			mockFENService.AssertNumberOfCalls(t, "Validate", 1)
		})
	}
}
