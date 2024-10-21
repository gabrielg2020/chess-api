package entity

import ("errors")

type MoveEntityInterface interface {
	GetFromX() (int, error)
	GetFromY() (int, error)
	GetToX() (int, error)
	GetToY() (int, error)
	GetPromotion() (int, error)
	IsCastling() (bool, error)
	IsEnPassant() (bool, error)
	GetCaptured() (int, error)
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

func (entity *MoveEntity) GetFromX() (int, error) {
	if entity.fromX == nil {
		return 0, errors.New("move.fromX is not set")
	}
	return *entity.fromX, nil
}

func (entity *MoveEntity) GetFromY() (int, error) {
	if entity.fromY == nil {
		return 0, errors.New("move.fromY is not set")
	}
	return *entity.fromY, nil
}

func (entity *MoveEntity) GetToX() (int, error) {
	if entity.toX == nil {
		return 0, errors.New("move.toX is not set")
	}
	return *entity.toX, nil
}

func (entity *MoveEntity) GetToY() (int, error) {
	if entity.toY == nil {
		return 0, errors.New("move.toY is not set")
	}
	return *entity.toY, nil
}

func (entity *MoveEntity) GetPromotion() (int, error) {
	if entity.promotion == nil {
		return 0, errors.New("move.promotion is not set")
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
		return 0, errors.New("move.captured is not set")
	}
	return *entity.captured, nil
}