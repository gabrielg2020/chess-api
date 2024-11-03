package entity

import (
	"errors"
	"github.com/gabrielg2020/chess-api/api/service/helper_service"
	"github.com/gabrielg2020/chess-api/pkg/logger"
	"github.com/sirupsen/logrus"
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
	GetChessNotation() (string, error)
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
	logger.Log.WithFields(logrus.Fields{
		"fromX": HelperService.IntValue(fromX), "fromY": HelperService.IntValue(fromY),
		"toX": HelperService.IntValue(toX), "toY": HelperService.IntValue(toY),
		"promotion":  HelperService.IntValue(promotion),
		"isCastling": HelperService.BoolValue(isCastling), "isEnPassant": HelperService.BoolValue(isEnPassant),
		"captured": HelperService.IntValue(captured),
	}).Debug("NewMoveEntity: Created a new MoveEntity")

	return &MoveEntity{
		fromX:       fromX,
		fromY:       fromY,
		toX:         toX,
		toY:         toY,
		promotion:   promotion,
		isCastling:  isCastling,
		isEnPassant: isEnPassant,
		captured:    captured,
	}
}

// Methods

func (entity *MoveEntity) GetFromX() (int, error) {
	if entity.fromX == nil {
		return -1, errors.New("MoveEntity.GetFromX: fromX is not set")
	}
	return *entity.fromX, nil
}

func (entity *MoveEntity) GetFromY() (int, error) {
	if entity.fromY == nil {
		return -1, errors.New("MoveEntity.GetFromY: fromY is not set")
	}
	return *entity.fromY, nil
}

func (entity *MoveEntity) GetToX() (int, error) {
	if entity.toX == nil {
		return -1, errors.New("MoveEntity.GetToX: toX is not set")
	}
	return *entity.toX, nil
}

func (entity *MoveEntity) GetToY() (int, error) {
	if entity.toY == nil {
		return -1, errors.New("MoveEntity.GetToY: toY is not set")
	}
	return *entity.toY, nil
}

func (entity *MoveEntity) GetPromotion() (int, error) {
	if entity.promotion == nil {
		return -1, errors.New("MoveEntity.GetPromotion: promotion is not set")
	}
	return *entity.promotion, nil
}

func (entity *MoveEntity) IsCastling() (bool, error) {
	if entity.isCastling == nil {
		return false, errors.New("MoveEntity.IsCastling: isCastling is not set")
	}
	return *entity.isCastling, nil
}

func (entity *MoveEntity) IsEnPassant() (bool, error) {
	if entity.isEnPassant == nil {
		return false, errors.New("MoveEntity.IsEnPassant: isEnPassant is not set")
	}
	return *entity.isEnPassant, nil
}

func (entity *MoveEntity) GetCaptured() (int, error) {
	if entity.captured == nil {
		return -1, errors.New("MoveEntity.GetCaptured: captured is not set")
	}
	return *entity.captured, nil
}

// TODO needs testing :(
func (entity *MoveEntity) GetChessNotation() (string, error) {
	fromX, err := entity.GetFromX()
	if err != nil {
		return "", errors.New("MoveEntity.GetChessNotation: " + err.Error())
	}
	fromY, err := entity.GetFromY()
	if err != nil {
		return "", errors.New("MoveEntity.GetChessNotation: " + err.Error())
	}
	toX, err := entity.GetToX()
	if err != nil {
		return "", errors.New("MoveEntity.GetChessNotation: " + err.Error())
	}
	toY, err := entity.GetToY()
	if err != nil {
		return "", errors.New("MoveEntity.GetChessNotation: " + err.Error())
	}

	fromFile := string(rune('a' + fromX))
	toFile := string(rune('a' + toX))

	// We 'flip' the board, because board[0][0] represents top left, not bottom left like on a normal chessboard
	fromRank := 8 - fromY
	toRank := 8 - toY

	return fromFile + strconv.Itoa(fromRank) + toFile + strconv.Itoa(toRank), nil
}
