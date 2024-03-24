package core

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewGame_InitialisesGrid(t *testing.T) {
	t.Parallel()
	// Act
	game := NewGame(20, 10)

	// Assert
	assert.Len(t, game.Grid, 10)
	assert.Len(t, game.Grid[0], 20)
}

func TestGame_Move_MovesToExpectedLocation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		performMoves    func(*Game)
		expectedPlayerX int
		expectedPlayerY int
	}{
		"D-R-U-L": {
			performMoves: func(game *Game) {
				game.MoveDown()
				game.MoveRight()
				game.MoveUp()
				game.MoveLeft()
			},
			expectedPlayerX: 0,
			expectedPlayerY: 0,
		},
		"D-D-R-R": {
			performMoves: func(game *Game) {
				game.MoveDown()
				game.MoveDown()
				game.MoveRight()
				game.MoveRight()
			},
			expectedPlayerX: 2,
			expectedPlayerY: 2,
		},
		"D-D-D-D-D-D-U-U": {
			performMoves: func(game *Game) {
				game.MoveDown()
				game.MoveDown()
				game.MoveDown()
				game.MoveDown()
				game.MoveDown()
				game.MoveDown()
				game.MoveUp()
				game.MoveUp()
			},
			expectedPlayerX: 0,
			expectedPlayerY: 4,
		},
		"R-R-R-R-R-L": {
			performMoves: func(game *Game) {
				game.MoveRight()
				game.MoveRight()
				game.MoveRight()
				game.MoveRight()
				game.MoveRight()
				game.MoveLeft()
			},
			expectedPlayerX: 4,
			expectedPlayerY: 0,
		},
	}

	for name, testData := range tests {
		testData := testData
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			game := NewGame(10, 10)

			// Act
			testData.performMoves(game)

			// Assert
			assert.Equal(t, testData.expectedPlayerX, game.PlayerX)
			assert.Equal(t, testData.expectedPlayerY, game.PlayerY)
		})
	}
}

func TestGame_MoveUp_DoesNotCrossBorder(t *testing.T) {
	t.Parallel()
	// Arrange
	game := NewGame(10, 10)

	// Act
	result := func() { game.MoveUp() }

	// Assert
	assert.NotPanics(t, result)
	assert.Equal(t, 0, game.PlayerY)
}

func TestGame_MoveLeft_DoesNotCrossBorder(t *testing.T) {
	t.Parallel()
	// Arrange
	game := NewGame(10, 10)

	// Act
	result := func() { game.MoveLeft() }

	// Assert
	assert.NotPanics(t, result)
	assert.Equal(t, 0, game.PlayerX)
}

func TestGame_MoveDown_DoesNotCrossBorder(t *testing.T) {
	t.Parallel()
	// Arrange
	game := NewGame(10, 10)
	game.PlayerY = 9

	// Act
	result := func() { game.MoveDown() }

	// Assert
	assert.NotPanics(t, result)
	assert.Equal(t, 9, game.PlayerY)
}

func TestGame_MoveRight_DoesNotCrossBorder(t *testing.T) {
	t.Parallel()
	// Arrange
	game := NewGame(10, 10)
	game.PlayerX = 9

	// Act
	result := func() { game.MoveRight() }

	// Assert
	assert.NotPanics(t, result)
	assert.Equal(t, 9, game.PlayerX)
}
