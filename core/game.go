package core

import "math/rand"

var random = rand.New(rand.NewSource(10))

func NewGame(width int, height int) *Game {
	grid := make([][]Tile, height)
	for row := range grid {
		grid[row] = make([]Tile, width)
	}

	game := &Game{Grid: grid}
	game.spawnApple()
	return game
}

type Game struct {
	Grid    [][]Tile
	PlayerY int
	PlayerX int

	Score int
}

func (g *Game) MoveUp() {
	if g.PlayerY == 0 {
		return
	}

	g.PlayerY--

	g.checkApple()
}

func (g *Game) MoveDown() {
	if g.PlayerY >= len(g.Grid)-1 {
		return
	}

	g.PlayerY++

	g.checkApple()
}

func (g *Game) MoveLeft() {
	if g.PlayerX == 0 {
		return
	}

	g.PlayerX--

	g.checkApple()
}

func (g *Game) MoveRight() {
	if g.PlayerX >= len(g.Grid[0])-1 {
		return
	}

	g.PlayerX++

	g.checkApple()
}

func (g *Game) checkApple() {
	if !g.Grid[g.PlayerY][g.PlayerX].HasApple {
		return
	}

	g.Grid[g.PlayerY][g.PlayerX].HasApple = false
	g.Score++

	g.spawnApple()
}

func (g *Game) spawnApple() {
	randomX := random.Int() % len(g.Grid[0])
	randomY := random.Int() % len(g.Grid)

	if randomX == g.PlayerX && randomY == g.PlayerY {
		g.spawnApple()
		return
	}

	if g.Grid[randomX][randomY].HasApple {
		g.spawnApple()
		return
	}

	g.Grid[randomX][randomY].HasApple = true
}

type Tile struct {
	HasApple bool
}
