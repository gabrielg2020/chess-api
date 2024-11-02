package FENService

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabrielg2020/chess-api/api/entity"
)

type FENServiceInterface interface {
	Validate(fen string) error
	Parse(validFen string) (entity.ChessboardEntityInterface, error)
	// ParseToBitboard(validFen string) (entity.ChessboardEntityInterface, error)
}

type FENService struct{}

func NewFENService() *FENService {
	return &FENService{}
}

// Methods

func (service *FENService) Validate(fen string) error {
	if fen == "" {
		return errors.New("FENService.Validate: FEN string empty")
	}

	fenRegex := `^(([rnbqkpRNBQKP1-8]{1,8}/){7}[rnbqkpRNBQKP1-8]{1,8})\s([wb])\s(-|[KQkq]{1,4})\s(-|[a-h][36])\s(\d+)\s(\d+)$`
	matched, err := regexp.MatchString(fenRegex, fen)

	if err != nil {
		return errors.New("FENService.Validate: error validating FEN string. you most likely edited the RegEx string... ")
	}
	if !matched {
		return errors.New("FENService.Validate: string is not a FEN")
	}

	return nil
}

func (service *FENService) Parse(validFen string) (entity.ChessboardEntityInterface, error) {
	// Split fen string and assign into seperate variables
	// REFRENCE: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
	fenParts := strings.Fields(strings.TrimSpace(validFen))
	if len(fenParts) != 6 {
		return nil, errors.New("FENService.Validate: expected 6 fields in fenParts")
	}

	piecePlacement, activeColour, castlingRights, enPassantSquare, halfmoveClock, fullmoveNumber := fenParts[0], fenParts[1], fenParts[2], fenParts[3], fenParts[4], fenParts[5]

	pieceToIntMap := map[string]int{
		"p": -1, "P": 1,
		"n": -2, "N": 2,
		"b": -3, "B": 3,
		"r": -4, "R": 4,
		"q": -5, "Q": 5,
		"k": -6, "K": 6,
	}

	rows := strings.Split(piecePlacement, "/")
	if len(rows) != 8 {
		return nil, errors.New("FENService.Validate: expected 8 rows in piece placement")
	}

	board := [8][8]int{}

	for i := 0; i < 8; i++ {
		row := rows[i]
		col := 0
		for j := 0; j < len(row); j++ {
			piece := string(row[j])

			if pieceAsInt, exists := pieceToIntMap[piece]; exists { // Add piece to board
				if col >= 8 {
					return nil, errors.New("FENService.Validate: too many pieces in row")
				}
				board[i][col] = pieceAsInt
				col++
			} else { // Add spaces to board
				emptySquares, err := strconv.Atoi(piece)
				if err != nil {
					return nil, errors.New("FENService.Validate: invalid character in row")
				}
				if (col + emptySquares) > 8 {
					return nil, errors.New("FENService.Validate: too many squares in row")
				}
				for k := 0; k < emptySquares; k++ {
					board[i][col] = 0
					col++
				}
			}
		}
		if col != 8 {
			return nil, errors.New("FENService.Validate: row does not have exactly 8 squares")
		}
	}

	chessboard := entity.NewChessboardEntity(
		&board,
		&validFen,
		&activeColour,
		&castlingRights,
		&enPassantSquare,
		&halfmoveClock,
		&fullmoveNumber,
	)

	return chessboard, nil
}
