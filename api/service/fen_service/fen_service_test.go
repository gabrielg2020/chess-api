package FENService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FENService_Validate(t *testing.T) {
	testCases := []struct {
		name          string
		fen           string
		expectedError error
	}{
		{
			name:          "Empty FEN",
			fen:           "",
			expectedError: errors.New("FEN string empty"),
		},
		{
			name:          "Valid FEN",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Black to play first]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [White already castled]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w kq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [White can castle King side]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [White can castle Queen side]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qkq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Black already castled]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Black can castle King side]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQk - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Black can castle Queen side]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQq - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Both players castled]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			expectedError: nil,
		},
		{
			name:          "Valid FEN [Possible en passant]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq e3 0 1",
			expectedError: nil,
		},
		{
			name:          "Invalid FEN",
			fen:           "invalid_fen",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove a row]",
			fen:           "pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove colour to play]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR  KQkq - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove castling rights]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w  - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove en passant]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq  0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove full turns]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -  1",
			expectedError: errors.New("string is not a FEN"),
		},
		{
			name:          "Invalid FEN [Remove half turns]",
			fen:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 ",
			expectedError: errors.New("string is not a FEN"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			service := NewFENService()
			// Act
			err := service.Validate(tc.fen)
			// Assert
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_FENService_Parse(t *testing.T) {
	testCases := []struct {
		name                    string
		fen                     string
		expectedBoard           [8][8]int
		expectedActiveColour    string
		expectedCastlingRights  string
		expectedEnPassantSquare string
		expectedHalfmoveClock   string
		expectedFullmoveNumber  string
		expectedError           error
	}{
		{
			name: "Starting Board",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedBoard: [8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			},
			expectedActiveColour:    "w",
			expectedCastlingRights:  "KQkq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock:   "0",
			expectedFullmoveNumber:  "1",
			expectedError:           nil,
		},
		{
			name:          "Invalid FEN - Empty String",
			fen:           "",
			expectedError: errors.New("expected 6 fields in fenParts"),
		},
		{
			name:          "Invalid FEN - Not Enough Fields",
			fen:           "8/8/8/8/8/8/8/8 w",
			expectedError: errors.New("expected 6 fields in fenParts"),
		},
		{
			name:          "Invalid FEN - Too Many Fields",
			fen:           "8/8/8/8/8/8/8/8 w KQkq - 0 1 extra",
			expectedError: errors.New("expected 6 fields in fenParts"),
		},
		{
			name:          "Invalid Piece Placement - Too Many Pieces in Row",
			fen:           "9/8/8/8/8/8/8/8 w KQkq - 0 1",
			expectedError: errors.New("too many squares in row"),
		},
		{
			name:          "Invalid Character in Piece Placement",
			fen:           "8/8/8/8/8/8/8/8x w KQkq - 0 1",
			expectedError: errors.New("invalid character in row"),
		},
		{
			name:                    "Empty Board",
			fen:                     "8/8/8/8/8/8/8/8 w - - 0 1",
			expectedBoard:           [8][8]int{},
			expectedActiveColour:    "w",
			expectedCastlingRights:  "-",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock:   "0",
			expectedFullmoveNumber:  "1",
			expectedError:           nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			service := NewFENService()

			// Act
			chessboard, err := service.Parse(tc.fen)

			if tc.expectedError != nil {
				assert.Nil(t, chessboard)
				assert.EqualError(t, err, tc.expectedError.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, chessboard)

			// Act
			board, boardErr := chessboard.GetBoard()
			activeColour, activeColourErr := chessboard.GetActiveColour()
			castlingRights, castlingRightsErr := chessboard.GetCastlingRights()
			enPassantSquare, enPassantSquareErr := chessboard.GetEnPassantSquare()
			halfmoveClock, halfmoveClockErr := chessboard.GetHalfmoveClock()
			fullmoveNumber, fullmoveNumberErr := chessboard.GetFullmoveNumber()

			// Assert: No errors from getters
			assert.NoError(t, boardErr)
			assert.NoError(t, activeColourErr)
			assert.NoError(t, castlingRightsErr)
			assert.NoError(t, enPassantSquareErr)
			assert.NoError(t, halfmoveClockErr)
			assert.NoError(t, fullmoveNumberErr)

			// Assert: Field values match expectations
			assert.Equal(t, tc.expectedBoard, board)
			assert.Equal(t, tc.expectedActiveColour, activeColour)
			assert.Equal(t, tc.expectedCastlingRights, castlingRights)
			assert.Equal(t, tc.expectedEnPassantSquare, enPassantSquare)
			assert.Equal(t, tc.expectedHalfmoveClock, halfmoveClock)
			assert.Equal(t, tc.expectedFullmoveNumber, fullmoveNumber)
		})
	}
}
