package game

import (
	"errors"
	"math"
)

type State string

const (
	STATE_RUNNING  State = "running"
	STATE_FINISHED State = "finished"
)

type Game struct {
	state   State
	turn    int
	players []*Player
	current *Player
	actions []Action
	score   []int
}

func CreateGame(p1, p2 *Player) *Game {
	players := make([]*Player, 2)
	players[0] = p1
	players[1] = p2

	actions := make([]Action, 9)
	score := make([]int, 8)

	return &Game{STATE_RUNNING, 0, players, p1, actions, score}
}

func (g *Game) hasWon() bool {
	for _, score := range g.score {
		if math.Abs(float64(score)) == 3 {
			return true
		}
	}
	return false
}

func (g *Game) hasMove(position Vertex) bool {
	for _, action := range g.actions {
		if action.player != nil && action.position.Equals(position) {
			return true
		}
	}

	return false
}

func (g *Game) TakeTurn(v Vertex) error {
	if g.state != STATE_RUNNING {
		return errors.New("Game has been finished.")
	}

	if g.hasMove(v) {
		return errors.New("This move has been done.")
	}

	a := Action{g.current, v}
	g.actions[g.turn] = a
	g.turn++

	points := 1
	if g.current == g.players[1] {
		points = -1
	}

	g.score[v.Y] += points
	g.score[3+v.X] += points

	if v.X == v.Y {
		g.score[6] += points
	}

	if 2-v.X == v.Y {
		g.score[7] += points
	}

	if g.hasWon() {
		g.state = STATE_FINISHED
		return nil
	}

	if g.current == g.players[0] {
		g.current = g.players[1]
	} else {
		g.current = g.players[0]
	}

	return nil
}

func (g *Game) GetState() State {
	return g.state
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.current
}
