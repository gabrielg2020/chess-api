package entity

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type ChessboardEntityInterface interface {
	GetBoard() ([8][8]int, error)
	GetFen() (string, error)
	GetActiveColour() (string, error)
	GetCastlingRights() (string, error)
	GetEnPassantSquare() (string, error)
	GetHalfmoveClock() (string, error)
	GetFullmoveNumber() (string, error)
	GetPiece(int, int) (int, error)
	IsSquareEmpty(int, int) (bool, error)
	IsOpponent(int, int, int) (bool, error)
	IsWithinBounds(int, int) bool
	SetSquare(int, int, int) error
	// SetFen(fen string) (*ChessboardEntity, error)
	// ResetBoard() (*ChessboardEntity, error)
	// SetPiece(position string, peice int) error
	// MovePiece(from string, to string) error
	// IsCheckmate(colour string) error
	// IsStalemate(colour string) error
	// IsCastlingAllowed(color string, side string) error
	// HasLegalMoves(colour string) error
	// GetLegalMoves(colour string) []string
	// RemovePeice(position string) error
	// PromotePawn(position string) error
	// GetPlayerPieces(colour string) ([][]int, error)
}

type ChessboardEntity struct {
	// White = +, Black = -
	// Pawns = 1, Knights = 2, Bishops = 3, Rooks = 4, Queens = 5, King = 6
	// e.g a white rook on square c5 -> board[2][4] = 4
	board           *[8][8]int
	fen             *string
	activeColour    *string
	castlingRights  *string
	enPassantSquare *string
	halfmoveClock   *string
	fullmoveNumber  *string
}

func NewChessboardEntity(board *[8][8]int, fen *string, activeColour *string, castlingRights *string, enPassantSquare *string, halfmoveClock *string, fullmoveNumber *string) *ChessboardEntity {
	return &ChessboardEntity{
		board:           board,
		fen:             fen,
		activeColour:    activeColour,
		castlingRights:  castlingRights,
		enPassantSquare: enPassantSquare,
		halfmoveClock:   halfmoveClock,
		fullmoveNumber:  fullmoveNumber,
	}
}

// Methods

func (entity *ChessboardEntity) GetBoard() ([8][8]int, error) {
	if entity.board == nil {
		return [8][8]int{}, errors.New("ChessboardEntity.GetBoard: board is not set")
	}
	return *entity.board, nil
}

func (entity *ChessboardEntity) SetBoard(board *[8][8]int) {
	entity.board = (*[8][8]int)(board)
}

func (entity *ChessboardEntity) GetFen() (string, error) {
	if entity.fen == nil {
		return "", errors.New("ChessboardEntity.GetFen: fen is not set")
	}
	return *entity.fen, nil
}

func (entity *ChessboardEntity) GetActiveColour() (string, error) {
	if entity.activeColour == nil {
		return "", errors.New("ChessboardEntity.GetActiveColour: activeColour is not set")
	}
	return *entity.activeColour, nil
}

func (entity *ChessboardEntity) GetCastlingRights() (string, error) {
	if entity.castlingRights == nil {
		return "", errors.New("ChessboardEntity.GetCastlingRights: castlingRights is not set")
	}
	return *entity.castlingRights, nil
}

func (entity *ChessboardEntity) GetEnPassantSquare() (string, error) {
	if entity.enPassantSquare == nil {
		return "", errors.New("ChessboardEntity.GetEnPassantSquare: enPassantSquare is not set")
	}
	return *entity.enPassantSquare, nil
}

func (entity *ChessboardEntity) SetEnPassantSquare(enPassantSquare *string) {
	entity.enPassantSquare = (*string)(enPassantSquare)
}

func (entity *ChessboardEntity) GetHalfmoveClock() (string, error) {
	if entity.halfmoveClock == nil {
		return "", errors.New("ChessboardEntity.GetHalfmoveClock: halfmoveClock is not set")
	}
	return *entity.halfmoveClock, nil
}

func (entity *ChessboardEntity) GetFullmoveNumber() (string, error) {
	if entity.fullmoveNumber == nil {
		return "", errors.New("ChessboardEntity.GetFullmoveNumber: fullmoveNumber is not set")
	}
	return *entity.fullmoveNumber, nil
}

func (entity *ChessboardEntity) GetPiece(row int, col int) (int, error) {
	if !entity.IsWithinBounds(row, col) {
		return -7, errors.New("ChessboardEntity.GetPiece: row or col out of bounds")
	}

	if entity.board == nil {
		return -7, errors.New("ChessboardEntity.GetPiece: board is not set")
	}

	return entity.board[row][col], nil
}

func (entity *ChessboardEntity) IsSquareEmpty(row int, col int) (bool, error) {
	if !entity.IsWithinBounds(row, col) {
		return false, nil
	}

	if entity.board == nil {
		return false, errors.New("ChessboardEntity.IsSquareEmpty: board is not set")
	}

	if entity.board[row][col] == 0 {
		return true, nil
	}
	return false, nil
}

func (entity *ChessboardEntity) IsOpponent(piece int, row int, col int) (bool, error) {
	if !entity.IsWithinBounds(row, col) {
		return false, nil
	}

	if entity.board == nil {
		return false, errors.New("ChessboardEntity.IsOpponent: board is not set")
	}

	// Check En Passant
	if piece == 1 || piece == -1 {
		enPassantSquare, err := entity.GetEnPassantSquare()
		if err != nil {
			return false, errors.New("ChessboardEntity.IsOpponent: " + err.Error())
		}

		// En Passant is not set
		if enPassantSquare == "-" {
			return false, nil
		}

		enPassantRow, enPassantCol, err := entity.convertChessNotation(enPassantSquare)
		if err != nil {
			return false, errors.New("ChessboardEntity.IsOpponent: " + err.Error())
		}

		if enPassantRow == row && enPassantCol == col {
			return true, nil
		}
	}

	if entity.board[row][col] == 0 {
		return false, nil
	}

	if math.Signbit(float64(entity.board[row][col])) == math.Signbit(float64(piece)) {
		return false, nil
	}

	return true, nil
}

func (entity *ChessboardEntity) IsWithinBounds(toX int, toY int) bool {
	if (toX > 7) || (toX < 0) || (toY > 7) || (toY < 0) {
		return false
	} else {
		return true
	}
}

func (entity *ChessboardEntity) convertChessNotation(chessNotation string) (int, int, error) {
	// remove spaces
	chessNotation = strings.TrimSpace(chessNotation)

	if len(chessNotation) < 2 {
		return -7, -7, errors.New("ChessboardEntity.convertChessNotation: invalid chess notation")
	}

	var letter, digit string

	// separate chess notation
	for _, char := range chessNotation {
		if unicode.IsLetter(char) {
			letter += string(char)
		} else if unicode.IsDigit(char) {
			digit += string(char)
		} else {
			return -7, -7, errors.New("ChessboardEntity.convertChessNotation: invalid character in chess notation")
		}
	}

	if len(letter) != 1 || len(digit) != 1 {
		return -7, -7, errors.New("ChessboardEntity.convertChessNotation: invalid chess notation format")
	}

	rowLetter := unicode.ToLower(rune(letter[0]))

	if rowLetter < 'a' || rowLetter > 'h' {
		return -7, -7, errors.New("ChessboardEntity.convertChessNotation: invalid column letter")
	}
	row := int(rowLetter - 'a')

	colNumber, err := strconv.Atoi(digit)
	if err != nil {
		return -7, -7, errors.New("ChessboardEntity.convertChessNotation: invalid row number")
	}
	if colNumber < 1 || colNumber > 8 {
		return -7, -7, errors.New("ChessboardEntity.convertChessNotation: row number out of range")
	}

	col := 8 - colNumber

	return col, row, nil
}

// TODO needs testing :(
func (entity *ChessboardEntity) SetSquare(row int, col int, piece int) error {
	if !entity.IsWithinBounds(row, col) {
		return errors.New("ChessboardEntity.SetSquare: row or col out of bounds")
	}

	if entity.board == nil {
		return errors.New("ChessboardEntity.SetSquare: board is not set")
	}

	entity.board[row][col] = piece

	return nil
}
