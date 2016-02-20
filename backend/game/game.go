package game

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"math"
)

type State string

const (
	STATE_RUNNING  State = "running"
	STATE_FINISHED State = "finished"
)

type Game struct {
	Id        string `gorethink:"id,omitempty"`
	Status    State
	Turn      int
	PlayerIds []string `gorethink:"players"`
	players   []*Player
	CurrentId string `gorethink:"current"`
	current   *Player
	Actions   []Action
	Score     []int
}

func RunningGameExists(p1, p2 *Player) (bool, error) {
	res, err := r.Table("game").Filter(r.Row.Field("Status").Eq(STATE_RUNNING)).Filter(func(g r.Term) r.Term {
		return g.Field("players").Contains(p1.Id, p2.Id)
	}).Run(session)

	if err != nil {
		return false, err
	}

	return !res.IsNil(), nil
}

func FindGame(p1, p2 *Player) (*Game, error) {
	res, err := r.Table("game").Filter(func(g r.Term) r.Term {
		return g.Field("players").Contains(p1.Id, p2.Id)
	}).Run(session)

	if err != nil {
		return nil, err
	}

	if res.IsNil() {
		return nil, nil
	}

	var g Game
	res.One(&g)

	g.players = make([]*Player, 2)

	for i, id := range g.PlayerIds {
		p, err := LoadPlayerByID(id)

		if err != nil {
			return nil, err
		}

		g.players[i] = p

		if g.CurrentId == id {
			g.current = p
		}
	}

	return &g, nil
}

func CreateGame(p1, p2 *Player) (*Game, error) {
	exists, err := RunningGameExists(p1, p2)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("There is already a game running.")
	}

	players := make([]*Player, 2)
	players[0] = p1
	players[1] = p2

	actions := make([]Action, 9)
	score := make([]int, 8)

	player_ids := []string{p1.Id, p2.Id}

	g := &Game{
		Status:    STATE_RUNNING,
		Turn:      0,
		PlayerIds: player_ids,
		players:   players,
		CurrentId: p1.Id,
		current:   p1,
		Actions:   actions,
		Score:     score,
	}

	_, err = r.Table("game").Insert(g).RunWrite(session)

	if err != nil {
		return nil, err
	}

	g, err = FindGame(p1, p2)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) Delete() error {
	_, err := r.Table("game").Get(g.Id).Delete().RunWrite(session)
	return err
}

func (g *Game) hasWon() bool {
	for _, score := range g.Score {
		if math.Abs(float64(score)) == 3 {
			return true
		}
	}
	return false
}

func (g *Game) hasMove(position Vertex) bool {
	for _, action := range g.Actions {
		if action.player != nil && action.position.Equals(position) {
			return true
		}
	}

	return false
}

func (g *Game) TakeTurn(v Vertex) error {
	if g.Status != STATE_RUNNING {
		return errors.New("Game has been finished.")
	}

	if g.hasMove(v) {
		return errors.New("This move has been done.")
	}

	a := Action{g.current, v}
	g.Actions[g.Turn] = a
	g.Turn++

	points := 1
	if g.current == g.players[1] {
		points = -1
	}

	g.Score[v.Y] += points
	g.Score[3+v.X] += points

	if v.X == v.Y {
		g.Score[6] += points
	}

	if 2-v.X == v.Y {
		g.Score[7] += points
	}

	if g.hasWon() {
		g.Status = STATE_FINISHED
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
	return g.Status
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.current
}
