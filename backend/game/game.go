package game

type State string

const (
	STATE_RUNNING State = "running"
	STATE_FINISHED = "finished"
)

type game struct  {
	state State
	turn int
}
