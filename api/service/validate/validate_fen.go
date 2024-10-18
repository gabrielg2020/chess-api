package validate

import (
	"errors"
	"regexp"
)

type FENServiceInterface interface {
	Validate(fen string) (bool, error)
}

type FENService struct {}

func NewFENService() *FENService {
	return &FENService{}
}

func (service *FENService) Validate(fen string) (bool, error) {
	if fen == "" {
		return false, errors.New("FEN string empty")
	}

	fenRegex := `^(([rnbqkpRNBQKP1-8]{1,8}/){7}[rnbqkpRNBQKP1-8]{1,8})\s([wb])\s(-|[KQkq]{1,4})\s(-|[a-h][36])\s(\d+)\s(\d+)$`
	matched, err := regexp.MatchString(fenRegex, fen)
	
	if err != nil {
		return false, errors.New("error validating FEN string. you most likely edited the RegEx string... ")
	}
	if !matched {
		return false, errors.New("string is not a FEN")
	}

	return true, nil
}
