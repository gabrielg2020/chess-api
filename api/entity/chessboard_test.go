package entity

import (
	"errors"
	"testing"

	HelperService "github.com/gabrielg2020/chess-api/api/service/helper_service"
	"github.com/stretchr/testify/assert"
)

func Test_ChessboardEntity_GetBoard(t *testing.T) {
	testCases := []struct {
		name             string
		board            *[8][8]int
		expectedResponse [8][8]int
		expectedError    error
	}{
		{
			name: "chessboard.board is set",
			board: &[8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			},
			expectedResponse: [8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			},
			expectedError: nil,
		},
		{
			name:             "chessboard.board is not set",
			board:            nil,
			expectedResponse: [8][8]int{},
			expectedError:    errors.New("ChessboardEntity.GetBoard: board is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(tc.board, nil, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetBoard()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetFen(t *testing.T) {
	testCases := []struct {
		name             string
		fen              *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.fen is set",
			fen:              HelperService.StrPtr("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			expectedResponse: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError:    nil,
		},
		{
			name:             "chessboard.fen is not set",
			fen:              nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetFen: fen is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, tc.fen, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetFen()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetActiveColour(t *testing.T) {
	testCases := []struct {
		name             string
		activeColour     *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.activeColour is set",
			activeColour:     HelperService.StrPtr("w"),
			expectedResponse: "w",
			expectedError:    nil,
		},
		{
			name:             "chessboard.activeColour is not set",
			activeColour:     nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetActiveColour: activeColour is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, tc.activeColour, nil, nil, nil, nil)
			// Act
			response, err := entity.GetActiveColour()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetCastlingRights(t *testing.T) {
	testCases := []struct {
		name             string
		castlingRights   *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.castlingRights is set",
			castlingRights:   HelperService.StrPtr("KQkq"),
			expectedResponse: "KQkq",
			expectedError:    nil,
		},
		{
			name:             "chessboard.castlingRights is not set",
			castlingRights:   nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetCastlingRights: castlingRights is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, tc.castlingRights, nil, nil, nil)
			// Act
			response, err := entity.GetCastlingRights()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetEnPassantSquare(t *testing.T) {
	testCases := []struct {
		name             string
		enPassantSquare  *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.enPassantSquare is set",
			enPassantSquare:  HelperService.StrPtr("e3"),
			expectedResponse: "e3",
			expectedError:    nil,
		},
		{
			name:             "chessboard.enPassantSquare is not set",
			enPassantSquare:  nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetEnPassantSquare: enPassantSquare is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, tc.enPassantSquare, nil, nil)
			// Act
			response, err := entity.GetEnPassantSquare()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetHalfmoveClock(t *testing.T) {
	testCases := []struct {
		name             string
		halfmoveClock    *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.halfmoveClock is set",
			halfmoveClock:    HelperService.StrPtr("0"),
			expectedResponse: "0",
			expectedError:    nil,
		},
		{
			name:             "chessboard.halfmoveClock is not set",
			halfmoveClock:    nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetHalfmoveClock: halfmoveClock is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, tc.halfmoveClock, nil)
			// Act
			response, err := entity.GetHalfmoveClock()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetFullmoveNumber(t *testing.T) {
	testCases := []struct {
		name             string
		fullmoveNumber   *string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "chessboard.fullmoveNumber is set",
			fullmoveNumber:   HelperService.StrPtr("1"),
			expectedResponse: "1",
			expectedError:    nil,
		},
		{
			name:             "chessboard.fullmoveNumber is not set",
			fullmoveNumber:   nil,
			expectedResponse: "",
			expectedError:    errors.New("ChessboardEntity.GetFullmoveNumber: fullmoveNumber is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, tc.fullmoveNumber)
			// Act
			response, err := entity.GetFullmoveNumber()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetPiece(t *testing.T) {
	testCases := []struct {
		name             string
		row              int
		col              int
		boardToSet       *[8][8]int
		expectedResponse int
		expectedError    error
	}{
		{
			name: "Get Piece From Valid Square",
			row:  0,
			col:  3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			expectedResponse: -5,
			expectedError:    nil,
		},
		{
			name:             "Get Piece When Board Is Not Set",
			row:              0,
			col:              3,
			boardToSet:       nil,
			expectedResponse: -7,
			expectedError:    errors.New("ChessboardEntity.GetPiece: board is not set"),
		},
		{
			name:             "Get Piece When Indexing Out Of Bounds",
			row:              8,
			col:              8,
			boardToSet:       nil,
			expectedResponse: -7,
			expectedError:    errors.New("ChessboardEntity.GetPiece: row or col out of bounds"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, nil)
			entity.SetBoard(tc.boardToSet)
			// Act
			response, err := entity.GetPiece(tc.row, tc.col)
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_IsSquareEmpty(t *testing.T) {
	testCases := []struct {
		name             string
		row              int
		col              int
		boardToSet       *[8][8]int
		expectedResponse bool
		expectedError    error
	}{
		{
			name: "Check On Empty Square",
			row:  2,
			col:  3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name: "Check On Non-Empty Square",
			row:  0,
			col:  3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:             "Check When Board Is Not Set",
			row:              0,
			col:              3,
			boardToSet:       nil,
			expectedResponse: false,
			expectedError:    errors.New("ChessboardEntity.IsSquareEmpty: board is not set"),
		},
		{
			name: "Check When Indexing Out Of Bounds",
			row:  8,
			col:  8,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			expectedResponse: false,
			expectedError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, nil)
			entity.SetBoard(tc.boardToSet)
			// Act
			response, err := entity.IsSquareEmpty(tc.row, tc.col)
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_IsOpponent(t *testing.T) {
	testCases := []struct {
		name                 string
		piece                int
		row                  int
		col                  int
		boardToSet           *[8][8]int
		enPassantSquareToSet *string
		expectedResponse     bool
		expectedError        error
	}{
		{
			name:  "Check On Opponent Square",
			piece: 5,
			row:   0,
			col:   3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     true,
			expectedError:        nil,
		},
		{
			name:  "Check On Non-Opponent Square",
			piece: -5,
			row:   0,
			col:   3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     false,
			expectedError:        nil,
		},
		{
			name:  "Check On Empty Square",
			piece: -5,
			row:   2,
			col:   3,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     false,
			expectedError:        nil,
		},
		{
			name:                 "Check When Board Is Not Set",
			piece:                -5,
			row:                  0,
			col:                  3,
			boardToSet:           nil,
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     false,
			expectedError:        errors.New("ChessboardEntity.IsOpponent: board is not set"),
		},
		{
			name:  "Check When Indexing Out Of Bounds",
			piece: -5,
			row:   8,
			col:   8,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     false,
			expectedError:        nil,
		},
		{
			name:  "Check En Passant",
			piece: -1,
			row:   4,
			col:   4,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, 0, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, -1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("e4"),
			expectedResponse:     true,
			expectedError:        nil,
		},
		{
			name:  "Check En Passant When En Passant Is Not Set",
			piece: -1,
			row:   4,
			col:   4,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, 0, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, -1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: nil,
			expectedResponse:     false,
			expectedError:        errors.New("ChessboardEntity.IsOpponent: ChessboardEntity.GetEnPassantSquare: enPassantSquare is not set"),
		},
		{
			name:  "Check En Passant When There Is No En Passant Square",
			piece: -1,
			row:   4,
			col:   4,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, 0, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, -1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("-"),
			expectedResponse:     false,
			expectedError:        nil,
		},
		{
			name:  "Check En Passant When Failing to Convert En Passant Square",
			piece: -1,
			row:   4,
			col:   4,
			boardToSet: HelperService.IntBoardArrayPtr([8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, 0, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, -1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			}),
			enPassantSquareToSet: HelperService.StrPtr("e24"),
			expectedResponse:     false,
			expectedError:        errors.New("ChessboardEntity.IsOpponent: ChessboardEntity.convertChessNotation: invalid chess notation format"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, nil)
			entity.SetBoard(tc.boardToSet)
			entity.SetEnPassantSquare(tc.enPassantSquareToSet)
			// Act
			response, err := entity.IsOpponent(tc.piece, tc.row, tc.col)
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_IsWithinBounds(t *testing.T) {
	testCases := []struct {
		name             string
		toX              int
		toY              int
		expectedResponse bool
	}{
		{
			name:             "Check When Out Of Bounds",
			toX:              8,
			toY:              8,
			expectedResponse: false,
		},
		{
			name:             "Check When In Bounds",
			toX:              3,
			toY:              5,
			expectedResponse: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, nil)
			// Act
			response := entity.IsWithinBounds(tc.toX, tc.toY)
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func Test_ChessboardEntity_convertChessNotation(t *testing.T) {
	testCases := []struct {
		name          string
		chessNotation string
		expectedRow   int
		expectedCol   int
		expectedErr   error
	}{
		{
			name:          "Valid Chess Notation - a1",
			chessNotation: "a1",
			expectedRow:   7,
			expectedCol:   0,
			expectedErr:   nil,
		},
		{
			name:          "Valid Chess Notation - g6",
			chessNotation: "g6",
			expectedRow:   2,
			expectedCol:   6,
			expectedErr:   nil,
		},
		{
			name:          "Valid Chess Notation - e6",
			chessNotation: "  e6  ",
			expectedRow:   2,
			expectedCol:   4,
			expectedErr:   nil,
		},
		{
			name:          "Invalid Chess Notation - too short",
			chessNotation: "6",
			expectedRow:   -7,
			expectedCol:   -7,
			expectedErr:   errors.New("ChessboardEntity.convertChessNotation: invalid chess notation"),
		},
		{
			name:          "Invalid Chess Notation - invalid char",
			chessNotation: "g2.",
			expectedRow:   -7,
			expectedCol:   -7,
			expectedErr:   errors.New("ChessboardEntity.convertChessNotation: invalid character in chess notation"),
		},
		{
			name:          "Invalid Chess Notation - invalid format",
			chessNotation: "ff3",
			expectedRow:   -7,
			expectedCol:   -7,
			expectedErr:   errors.New("ChessboardEntity.convertChessNotation: invalid chess notation format"),
		},
		{
			name:          "Invalid Chess Notation - invalid letter",
			chessNotation: "z5",
			expectedRow:   -7,
			expectedCol:   -7,
			expectedErr:   errors.New("ChessboardEntity.convertChessNotation: invalid column letter"),
		},
		{
			name:          "Invalid Chess Notation - invalid number",
			chessNotation: "c9",
			expectedRow:   -7,
			expectedCol:   -7,
			expectedErr:   errors.New("ChessboardEntity.convertChessNotation: row number out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, nil)
			// Act
			row, col, err := entity.convertChessNotation(tc.chessNotation)
			// Assert
			if tc.expectedErr != nil {
				assert.EqualError(t, tc.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedRow, row)
			assert.Equal(t, tc.expectedCol, col)
		})
	}
}
