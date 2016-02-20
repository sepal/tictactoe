package game

import (
	"time"
)

type Player struct {
	id       int
	nickname string
}

type PlayerSession struct {
	SessionKey  string
	Player      *Player
	LoginInTime time.Time
	LastSeen    time.Time
}

func CreatePlayer(nickname string) *Player {
	return &Player{0, nickname}
}