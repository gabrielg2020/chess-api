package BestMoveHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking
// FEN Service
type mockFENService struct {
	mock.Mock
}

func (m *mockFENService) Validate(fen string) (error) {
	args := m.Called(fen)
	return args.Error(0)
}

func (m *mockFENService) Parse(validFen string) (entity.ChessboardEntityInterface, error) {
	args := m.Called(validFen)
	return args.Get(0).(entity.ChessboardEntityInterface), args.Error(1)
}

// Best Move Service
type mockBestMoveService struct {
	mock.Mock
}

func (m *mockBestMoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error) {
	args := m.Called(chessboard)
	return args.String(0), args.Error(1)
}

// Chessboard Entity
type mockChessboardEntity struct {
	mock.Mock
}

func (m *mockChessboardEntity) GetBoard() ([8][8]int, error) {
	args := m.Called()
	return args.Get(0).([8][8]int), args.Error(1)
}

func (m *mockChessboardEntity) GetFen() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetActiveColour() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetCastlingRights() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetEnPassantSquare() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetHalfmoveClock() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetFullmoveNumber() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func Test_BestMoveHandler_FindBestMove(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name                 string
		fen                  string
		setupFENService      func(m *mockFENService)
		setupBestMoveService func(m *mockBestMoveService)
		expectedStatusCode   int
		expectedResponse     string
	}{
		{
			name: "Test Case 1",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			setupFENService: func(m *mockFENService) {
				m.On("Validate", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(nil)
				mockChessboard := new(mockChessboardEntity)
				m.On("Parse", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").Return(mockChessboard, nil)
			},
			setupBestMoveService: func(m *mockBestMoveService) {
				m.On("FindBestMove", mock.Anything).Return("a2a4", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"bestMove":"a2a4"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockFENService := new(mockFENService)
			mockBestMoveService := new(mockBestMoveService)
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

