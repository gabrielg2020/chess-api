package entity

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
	fromX       int
	fromY       int
	toX         int
	toY         int 
	promotion   int // 0 if not a promotion, otherwise the piece code
	isCastling  bool
	isEnPassant bool
	captured    int // 0 if no piece was captured, otherwise hold the piece code of the piece captured
}

// TODO Create getter functions for all methods

func (entity *MoveEntity) GetFromX() (int, error) {
	return entity.fromX, nil
}

func (entity *MoveEntity) GetFromY() (int, error) {
	return entity.fromY, nil
}

func (entity *MoveEntity) GetToX() (int, error) {
	return entity.toX, nil
}

func (entity *MoveEntity) GetToY() (int, error) {
	return entity.toY, nil
}

func (entity *MoveEntity) GetPromotion() (int, error) {
	return entity.promotion, nil
}

func (entity *MoveEntity) IsCastling() (bool, error) {
	return entity.isCastling, nil
}

func (entity *MoveEntity) IsEnPassant() (bool, error) {
	return entity.isEnPassant, nil
}

func (entity *MoveEntity) GetCaptured() (int, error) {
  return entity.captured, nil
}