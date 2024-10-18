package validate

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFENService_Validate(t *testing.T) {
	// Arrange
	service := NewFENService()

	testCases := []struct {
		name    string
		fen     string
		isValid bool
		wantErr bool
	}{
		{ 
			name:    "Valid FEN",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Black to play first]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [White already castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w kq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [White can castle King side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [White can castle Queen side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qkq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Black already castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Black can castle King side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQk - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Black can castle Queen side]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQq - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Both players castled]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Valid FEN [Possible en passant]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq e3 0 1",
			isValid: true,
			wantErr: false,
		},
		{ 
			name:    "Invalid FEN",
			fen:     "invalid_fen",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Empty FEN",
			fen:     "",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove a row]",
			fen:     "pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove colour to play]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR  KQkq - 0 1",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove castling rights]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w  - 0 1",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove en passant]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq  0 1",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove full turns]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -  1",
			isValid: false,
			wantErr: true,
		},
		{ 
			name:    "Invalid FEN [Remove half turns]",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 ",
			isValid: false,
			wantErr: true,
		},		
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			isValid, _ := service.Validate(tc.fen)
			// Assert
			if tc.wantErr{
				assert.False(t, isValid, "Expected FEN to be invalid")
			} else {
				assert.True(t, isValid, "Expected FEN to be valid")
			}
		})
	}

}