package game

type State string

const (
	STATE_RUNNING State = "running"
	STATE_FINISHED = "finished"
)

type Game struct {
	state   State
	turn    int
	players [2]*Player
	current *Player
	actions []Action
}
