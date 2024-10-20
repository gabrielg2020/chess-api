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
			name:    "Empty FEN",
			fen:     "",
			expectedError: errors.New("FEN string empty"),
		},
		{ 
			name:    "Valid FEN",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Black to play first]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [White already castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w kq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [White can castle King side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [White can castle Queen side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qkq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Black already castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Black can castle King side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQk - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Black can castle Queen side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQq - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Both players castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Valid FEN [Possible en passant]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq e3 0 1",
			expectedError: nil,
		},
		{ 
			name:    "Invalid FEN",
			fen:     "invalid_fen",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove a row]",
			fen:     "pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove colour to play]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR  KQkq - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove castling rights]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w  - 0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove en passant]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq  0 1",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove full turns]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -  1",
			expectedError: errors.New("string is not a FEN"),
		},
		{ 
			name:    "Invalid FEN [Remove half turns]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 ",
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
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_FENService_Parse(t *testing.T){
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
	} {
		{
			name: "Starting Board",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
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
			expectedActiveColour: "w",
			expectedCastlingRights: "KQkq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "Custom Board",
			fen: "rnbqkbnr/ppp2ppp/8/4pPnN/8/bBrRqQkK/PPPPPPPP/3QKBNR w KQkq - 0 1",
			expectedBoard: [8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, 0, 0, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, -1, 1, -2, 2},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{-3, 3, -4, 4, -5, 5, -6, 6},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 5, 6, 3, 2, 4},
			},
			expectedActiveColour: "w",
			expectedCastlingRights: "KQkq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "Clear Board",
			fen: "PPPPPPPP/PPPPPPPP/PPPPPPPP/PPPPPPPP/PPPPPPPP/PPPPPPPP/PPPPPPPP/PPPPPPPP w KQkq - 0 1",
			expectedBoard: [8][8]int{
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
			},
			expectedActiveColour: "w",
			expectedCastlingRights: "KQkq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "Only white has Castling Rights",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1",
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
			expectedActiveColour: "w",
			expectedCastlingRights: "KQ",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "Only black has Castling Rights",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w kq - 0 1",
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
			expectedActiveColour: "w",
			expectedCastlingRights: "kq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "No Castling Rights",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
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
			expectedActiveColour: "w",
			expectedCastlingRights: "-",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "En Passant Square active",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq e3 0 1",
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
			expectedActiveColour: "w",
			expectedCastlingRights: "KQkq",
			expectedEnPassantSquare: "e3",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		{
			name: "Full Board",
			fen: "8/8/8/8/8/8/8/8 w KQkq - 0 1",
			expectedBoard: [8][8]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectedActiveColour: "w",
			expectedCastlingRights: "KQkq",
			expectedEnPassantSquare: "-",
			expectedHalfmoveClock: "0",
			expectedFullmoveNumber: "1",
			expectedError: nil,
		},
		
		{
			name: "Empty FEN",
			fen: "",
			expectedBoard: [8][8]int{},
			expectedError: errors.New("expected 6 feilds in fenParts"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
				// Arange
				service := NewFENService()
				// Act
				chessboard, err := service.Parse(tc.fen)
				board, _ := chessboard.GetBoard()
				activeColour, _ := chessboard.GetActiveColour()
				castlingRights, _ := chessboard.GetCastlingRights()
				enPassantSquare, _ := chessboard.GetEnPassantSquare()
				halfmoveClock, _ := chessboard.GetHalfmoveClock()
				fullmoveNumber, _ := chessboard.GetFullmoveNumber()

				// Assert
				// error
				assert.Equal(t, tc.expectedError, err)
				// values
				assert.Equal(t, tc.expectedBoard, board)
				assert.Equal(t, tc.expectedActiveColour, activeColour)
				assert.Equal(t, tc.expectedCastlingRights, castlingRights)
				assert.Equal(t, tc.expectedEnPassantSquare, enPassantSquare)
				assert.Equal(t, tc.expectedHalfmoveClock, halfmoveClock)
				assert.Equal(t, tc.expectedFullmoveNumber, fullmoveNumber)
		})
	}
}