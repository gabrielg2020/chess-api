package entity

import (
	"errors"
	"math"
	"strconv"
)

type MoveEntityInterface interface {
	GetFromX() (int, error)
	GetFromY() (int, error)
	GetToX() (int, error)
	GetToY() (int, error)
	GetPromotion() (int, error)
	IsCastling() (bool, error)
	IsEnPassant() (bool, error)
	GetCaptured() (int, error)
	GetChessNotation() (string)
}

type MoveEntity struct {
	fromX       *int
	fromY       *int
	toX         *int
	toY         *int
	promotion   *int // 0 if not a promotion, otherwise the piece code
	isCastling  *bool
	isEnPassant *bool
	captured    *int // 0 if no piece was captured, otherwise hold the piece code of the piece captured
} 

func NewMoveEntity(fromX *int, fromY *int, toX *int, toY *int, promotion *int, isCastling *bool, isEnPassant *bool, captured *int) *MoveEntity {
	return &MoveEntity{
		fromX: fromX,
		fromY: fromY,
		toX: toX,
		toY: toY,
		promotion: promotion,
		isCastling: isCastling,
		isEnPassant: isEnPassant,
		captured: captured,
	}
}

// Methods

func (entity *MoveEntity) GetFromX() (int, error) {
	if entity.fromX == nil {
		return -1, errors.New("move.fromX is not set")
	}
	return *entity.fromX, nil
}

func (entity *MoveEntity) GetFromY() (int, error) {
	if entity.fromY == nil {
		return -1, errors.New("move.fromY is not set")
	}
	return *entity.fromY, nil
}

func (entity *MoveEntity) GetToX() (int, error) {
	if entity.toX == nil {
		return -1, errors.New("move.toX is not set")
	}
	return *entity.toX, nil
}

func (entity *MoveEntity) GetToY() (int, error) {
	if entity.toY == nil {
		return -1, errors.New("move.toY is not set")
	}
	return *entity.toY, nil
}

func (entity *MoveEntity) GetPromotion() (int, error) {
	if entity.promotion == nil {
		return -1, errors.New("move.promotion is not set")
	}
	return *entity.promotion, nil
}

func (entity *MoveEntity) IsCastling() (bool, error) {
	if entity.isCastling == nil {
		return false, errors.New("move.isCastling is not set")
	}
	return *entity.isCastling, nil
}

func (entity *MoveEntity) IsEnPassant() (bool, error) {
	if entity.isEnPassant == nil {
		return false, errors.New("move.isEnPassant is not set")
	}
	return *entity.isEnPassant, nil
}

func (entity *MoveEntity) GetCaptured() (int, error) {
	if entity.captured == nil {
		return -1, errors.New("move.captured is not set")
	}
	return *entity.captured, nil
}

// TODO needs testing :(
// BUG may need a refactor to be more robust, take a look at chessboard.convertChessNotation
func (entity *MoveEntity) GetChessNotation() (string) {
	rowToChessNotationRow := map[int]string{
		0: "a", 1: "b",
		2: "c", 3: "d",
		4: "e", 5: "f",
		6: "g", 7: "h",
	}

	fromX, _ := entity.GetFromX()
	fromY, _ := entity.GetFromY()
	toX, _ := entity.GetToX()
	toY, _ := entity.GetToY()

	// We 'flip' the board, because board[0][0] represents top left, not bottom left like on a normal chessboard
	fromY = int(math.Abs(float64(fromY) - (8)))
	toY = int(math.Abs(float64(toY) - (8)))
	return rowToChessNotationRow[fromX] + strconv.FormatInt(int64(fromY), 10) + rowToChessNotationRow[toX] + strconv.FormatInt(int64(toY), 10)
}