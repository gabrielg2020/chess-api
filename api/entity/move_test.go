package entity

import (
	"errors"
	"testing"

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
			name: "move.fromX is set",
			fromX: func() *int {
				i := 3
				return &i
			}(),
			expectedResponse: 3,
			expectedError: nil,
		},
		{
			name: "move.fromX is not set",
			fromX: nil,
			expectedResponse: -1,
			expectedError: errors.New("move.fromX is not set"),
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
			name: "move.fromY is set",
			fromY: func() *int {
				i := 4
				return &i
			}(),
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
			name: "move.toX is set",
			toX: func() *int {
				i := 5
				return &i
			}(),
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
			name: "move.toY is set",
			toY: func() *int {
				i := 6
				return &i
			}(),
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
			name: "move.promotion is set",
			promotion: func() *int {
				i := 5 // Promoting to a Queen
				return &i
			}(),
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
			name: "move.isCastling is set to true",
			isCastling: func() *bool {
				b := true
				return &b
			}(),
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name: "move.isCastling is set to false",
			isCastling: func() *bool {
				b := false
				return &b
			}(),
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
			name: "move.isEnPassant is set to true",
			isEnPassant: func() *bool {
				b := true
				return &b
			}(),
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name: "move.isEnPassant is set to false",
			isEnPassant: func() *bool {
				b := false
				return &b
			}(),
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
			name: "move.captured is set",
			captured: func() *int {
				i := 2 // Captured a Knight
				return &i
			}(),
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