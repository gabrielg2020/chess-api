package FENService

import (
	"errors"
	"regexp"

	"github.com/gabrielg2020/chess-api/api/entity"
)

type FENServiceInterface interface {
	Validate(fen string) (error)
	Parse(validFen string) (entity.ChessboardEntityInterface, error)
	// ParseToBitboard(validFen string) (entity.ChessboardEntityInterface, error)
}

type FENService struct {}

func NewFENService() *FENService {
	return &FENService{}
}

// Methods

func (service *FENService) Validate(fen string) (error) {
	if fen == "" {
		return errors.New("FEN string empty")
	}

	fenRegex := `^(([rnbqkpRNBQKP1-8]{1,8}/){7}[rnbqkpRNBQKP1-8]{1,8})\s([wb])\s(-|[KQkq]{1,4})\s(-|[a-h][36])\s(\d+)\s(\d+)$`
	matched, err := regexp.MatchString(fenRegex, fen)
	
	if err != nil {
		return errors.New("error validating FEN string. you most likely edited the RegEx string... ")
	}
	if !matched {
		return errors.New("string is not a FEN")
	}

	return nil
}

func (service *FENService) Parse(validFen string) (entity.ChessboardEntityInterface, error) {
	// TODO Create parsing functionality
	var board [8][8]int
	chessboard := entity.NewChessboardEntity(board, validFen)
	return chessboard, nil
}
