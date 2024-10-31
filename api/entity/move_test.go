package entity

import (
	"errors"
	"testing"

	"github.com/gabrielg2020/chess-api/api/service/helper_service"
	"github.com/stretchr/testify/assert"
)

func Test_MoveEntity_GetFromX(t *testing.T) {
	testCase := []struct {
		name             string
		fromX            *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.fromX is set",
			fromX:            HelperService.IntPtr(3),
			expectedResponse: 3,
			expectedError:    nil,
		},
		{
			name:             "move.fromX is not set",
			fromX:            nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.fromX is not set"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(tc.fromX, nil, nil, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetFromX()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetFromY(t *testing.T) {
	testCases := []struct {
		name             string
		fromY            *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.fromY is set",
			fromY:            HelperService.IntPtr(4),
			expectedResponse: 4,
			expectedError:    nil,
		},
		{
			name:             "move.fromY is not set",
			fromY:            nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.fromY is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, tc.fromY, nil, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetFromY()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetToX(t *testing.T) {
	testCases := []struct {
		name             string
		toX              *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.toX is set",
			toX:              HelperService.IntPtr(5),
			expectedResponse: 5,
			expectedError:    nil,
		},
		{
			name:             "move.toX is not set",
			toX:              nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.toX is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, tc.toX, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetToX()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetToY(t *testing.T) {
	testCases := []struct {
		name             string
		toY              *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.toY is set",
			toY:              HelperService.IntPtr(6),
			expectedResponse: 6,
			expectedError:    nil,
		},
		{
			name:             "move.toY is not set",
			toY:              nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.toY is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, nil, tc.toY, nil, nil, nil, nil)
			// Act
			response, err := entity.GetToY()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetPromotion(t *testing.T) {
	testCases := []struct {
		name             string
		promotion        *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.promotion is set",
			promotion:        HelperService.IntPtr(5), // Promoting to a Queen
			expectedResponse: 5,
			expectedError:    nil,
		},
		{
			name:             "move.promotion is not set",
			promotion:        nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.promotion is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, nil, nil, tc.promotion, nil, nil, nil)
			// Act
			response, err := entity.GetPromotion()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_IsCastling(t *testing.T) {
	testCases := []struct {
		name             string
		isCastling       *bool
		expectedResponse bool
		expectedError    error
	}{
		{
			name:             "move.isCastling is set to true",
			isCastling:       HelperService.BoolPtr(true),
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name:             "move.isCastling is set to false",
			isCastling:       HelperService.BoolPtr(false),
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:             "move.isCastling is not set",
			isCastling:       nil,
			expectedResponse: false,
			expectedError:    errors.New("move.isCastling is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, nil, nil, nil, tc.isCastling, nil, nil)
			// Act
			response, err := entity.IsCastling()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_IsEnPassant(t *testing.T) {
	testCases := []struct {
		name             string
		isEnPassant      *bool
		expectedResponse bool
		expectedError    error
	}{
		{
			name:             "move.isEnPassant is set to true",
			isEnPassant:      HelperService.BoolPtr(true),
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name:             "move.isEnPassant is set to false",
			isEnPassant:      HelperService.BoolPtr(false),
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:             "move.isEnPassant is not set",
			isEnPassant:      nil,
			expectedResponse: false,
			expectedError:    errors.New("move.isEnPassant is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, nil, nil, nil, nil, tc.isEnPassant, nil)
			// Act
			response, err := entity.IsEnPassant()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetCaptured(t *testing.T) {
	testCases := []struct {
		name             string
		captured         *int
		expectedResponse int
		expectedError    error
	}{
		{
			name:             "move.captured is set",
			captured:         HelperService.IntPtr(2), // Captured a Knight
			expectedResponse: 2,
			expectedError:    nil,
		},
		{
			name:             "move.captured is not set",
			captured:         nil,
			expectedResponse: -1,
			expectedError:    errors.New("move.captured is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(nil, nil, nil, nil, nil, nil, nil, tc.captured)
			// Act
			response, err := entity.GetCaptured()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_MoveEntity_GetChessNotation(t *testing.T) {
	testCases := []struct {
		name             string
		fromX            *int
		fromY            *int
		toX              *int
		toY              *int
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "Valid move.fromX, move.fromY, move.toX, move.toY",
			fromX:            HelperService.IntPtr(1),
			fromY:            HelperService.IntPtr(2),
			toX:              HelperService.IntPtr(3),
			toY:              HelperService.IntPtr(4),
			expectedResponse: "b6d4",
			expectedError:    nil,
		},
		{
			name:             "Invalid move.fromX",
			fromX:            nil,
			fromY:            HelperService.IntPtr(2),
			toX:              HelperService.IntPtr(3),
			toY:              HelperService.IntPtr(4),
			expectedResponse: "",
			expectedError:    errors.New("failed to get fromX"),
		},
		{
			name:             "Invalid move.fromY",
			fromX:            HelperService.IntPtr(1),
			fromY:            nil,
			toX:              HelperService.IntPtr(3),
			toY:              HelperService.IntPtr(4),
			expectedResponse: "",
			expectedError:    errors.New("failed to get fromY"),
		},
		{
			name:             "Invalid move.toX",
			fromX:            HelperService.IntPtr(1),
			fromY:            HelperService.IntPtr(2),
			toX:              nil,
			toY:              HelperService.IntPtr(4),
			expectedResponse: "",
			expectedError:    errors.New("failed to get toX"),
		},
		{
			name:             "Invalid move.toY",
			fromX:            HelperService.IntPtr(1),
			fromY:            HelperService.IntPtr(2),
			toX:              HelperService.IntPtr(4),
			toY:              nil,
			expectedResponse: "",
			expectedError:    errors.New("failed to get toY"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewMoveEntity(tc.fromX, tc.fromY, tc.toX, tc.toY, nil, nil, nil, nil)
			// Act
			response, err := entity.GetChessNotation()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
