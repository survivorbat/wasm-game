//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"github.com/survivorbat/wasm-game/core"
	"syscall/js"
)

var game *core.Game

func main() {
	fmt.Println("Starting new game")
	game = core.NewGame(10, 10)

	// Register JS calls
	js.Global().Set("moveLeft", js.FuncOf(moveLeft))
	js.Global().Set("moveRight", js.FuncOf(moveRight))
	js.Global().Set("moveUp", js.FuncOf(moveUp))
	js.Global().Set("moveDown", js.FuncOf(moveDown))
	js.Global().Set("gameState", js.FuncOf(gameState))

	// Hang forever
	select {}
}

func moveLeft(js.Value, []js.Value) any {
	game.MoveLeft()
	return nil
}

func moveUp(js.Value, []js.Value) any {
	game.MoveUp()
	return nil
}

func moveDown(js.Value, []js.Value) any {
	game.MoveDown()
	return nil
}

func moveRight(js.Value, []js.Value) any {
	game.MoveRight()
	return nil
}

func gameState(js.Value, []js.Value) any {
	state, _ := json.Marshal(game)
	return string(state)
}
