package entity

type Move struct {
	fromX       int
	fromY       int
	toX         int
	toY         int 
	promotion   int // 0 if not a promotion, otherwise the piece code
	isCastling  bool
	isEnPassant bool
	captured    int // 0 if no piece was captured, otherwise hold the piece code of the piece captured
}