package entity

import (
	"errors"
)

type ChessboardEntityInterface interface {
	GetBoard() ([8][8]int, error)
	GetFen() (string, error)
	GetActiveColour() (string, error)
	GetCastlingRights() (string, error)
	GetEnPassantSquare() (string, error)
	GetHalfmoveClock() (string, error)
	GetFullmoveNumber() (string, error)
	// SetFen(fen string) (*ChessboardEntity, error)
	// ResetBoard() (*ChessboardEntity, error)
	// GetPiece(position string) (int, error)
	// SetPiece(position string, peice int) error
	// MovePiece(from string, to string) error
	// IsSquareEmpty(position string) error
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
	board           [8][8]int
	fen             string
	activeColour    string
	castlingRights  string
	enPassantSquare string
	halfmoveClock   string
	fullmoveNumber  string
}

func NewChessboardEntity(board [8][8]int, fen string, activeColour string, castlingRights string, enPassantSquare string, halfmoveClock string, fullmoveNumber string) *ChessboardEntity {
	return &ChessboardEntity{
		board: board,
		fen: fen,
		activeColour: activeColour,
		castlingRights: castlingRights,
		enPassantSquare: enPassantSquare,
		halfmoveClock: halfmoveClock,
		fullmoveNumber: fullmoveNumber,
	}
}

// Methods

func (entity *ChessboardEntity) GetBoard() ([8][8]int, error){
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if entity.board[row][col] > 6 || entity.board[row][col] < -6 {
				return [8][8]int{}, errors.New("chessboard.board is not set")
			}
		}
	}
	return entity.board, nil
}

func (entity *ChessboardEntity) GetFen() (string, error) {
	if entity.fen == "" {
		return "", errors.New("chessboard.fen is not set")
	}
	return entity.fen, nil
}

func (entity *ChessboardEntity) GetActiveColour() (string, error){
	if entity.activeColour == "" {
		return "", errors.New("chessboard.activeColour is not set")
	}
	return entity.activeColour, nil
}

func (entity *ChessboardEntity) GetCastlingRights() (string, error){
	if entity.castlingRights == "" {
		return "", errors.New("chessboard.castlingRights is not set")
	}
	return entity.castlingRights, nil
}

func (entity *ChessboardEntity) GetEnPassantSquare() (string, error){
	if entity.enPassantSquare == "" {
		return "", errors.New("chessboard.enPassantSquare is not set")
	}
	return entity.enPassantSquare, nil
}

func (entity *ChessboardEntity) GetHalfmoveClock() (string, error){
	if entity.halfmoveClock == "" {
		return "", errors.New("chessboard.halfmoveClock is not set")
	}
	return entity.halfmoveClock, nil
}

func (entity *ChessboardEntity) GetFullmoveNumber() (string, error){
	if entity.fullmoveNumber == "" {
		return "", errors.New("chessboard.fullmoveNumber is not set")
	}
	return entity.fullmoveNumber, nil
}
