package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Game struct {
	rolls []int
}

func (game *Game) Roll(pins int) {
	game.rolls = append(game.rolls, pins)
}

func (game *Game) Score() int {
	score := 0
	i := 0
	for nFrame := 0; nFrame < 10; nFrame++ {
		if isStrike(game, i) {
			score += 10 + game.rolls[i+1] + game.rolls[i+2]
			i += 1
		} else if isSpare(game, i) {
			score += 10 + game.rolls[i+2]
			i += 2
		} else {
			score += game.rolls[i] + game.rolls[i+1]
			i += 2
		}
	}
	return score
}

func isSpare(game *Game, i int) bool {
	return game.rolls[i]+game.rolls[i+1] == 10
}

func isStrike(game *Game, i int) bool {
	return game.rolls[i] == 10
}

func TestBowling(t *testing.T) {
	game := &Game{}
	manyRolls(game, 20, 0)
	assert.Equal(t, 0, game.Score())

	game = &Game{}
	manyRolls(game, 20, 1)
	assert.Equal(t, 20, game.Score())

	game = &Game{}
	manyRolls(game, 20, 2)
	assert.Equal(t, 40, game.Score())

	game = &Game{}
	game.Roll(5)
	game.Roll(5)
	manyRolls(game, 18, 1)
	assert.Equal(t, 29, game.Score())

	game = &Game{}
	manyRolls(game, 21, 5)
	assert.Equal(t, 150, game.Score())

	game = &Game{}
	game.Roll(10)
	manyRolls(game, 18, 1)
	assert.Equal(t, 30, game.Score())

	game = &Game{}
	manyRolls(game, 12, 10)
	assert.Equal(t, 300, game.Score())
}

func manyRolls(game *Game, times int, pins int) {
	for i := 0; i < times; i++ {
		game.Roll(pins)
	}
}
