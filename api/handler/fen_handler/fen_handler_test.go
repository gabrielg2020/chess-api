package FENHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"


	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock service
type mockFENService struct {
	mock.Mock
}

func (m *mockFENService) Validate(fen string) (bool, error) {
	args := m.Called(fen)
	return args.Bool(0), args.Error(1)
}

func TestValidateFEN(t *testing.T) {
	// Arange
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name               string
		fen                string
		mockIsValid        bool
		mockError          error
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "Valid FEN",
			fen:                "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			mockIsValid:        true,
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"valid":true}`,	
		},
		{
			name:               "Invalid FEN [FEN string doesn't pass regex]",
			fen:                "fen",
			mockIsValid:        false,
			mockError:          errors.New("string is not a FEN"),
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"valid":false, "errorMessage":"string is not a FEN", "errorCode":400}`,
		},
		{
			name:               "Invalid FEN [FEN string is empty]",
			fen:                "",
			mockIsValid:        false,
			mockError:          errors.New("FEN string empty"),
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"valid":false, "errorMessage":"FEN string empty", "errorCode":400}`,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFENService := new(mockFENService)
			handler := NewFENHandler(mockFENService)
		
			engine := gin.Default()
			engine.GET("/validate/fen", handler.ValidateFEN)

			mockFENService.On("Validate", tc.fen).Return(tc.mockIsValid, tc.mockError)

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