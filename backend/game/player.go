package game

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"time"
)

type Player struct {
	Id       string `gorethink:"id,omitempty"`
	Nickname string
	Created  time.Time
}

type PlayerSession struct {
	SessionKey  string
	Player      *Player
	LoginInTime time.Time
	LastSeen    time.Time
}

func NicknameExists(nickname string) (bool, error) {
	res, err := r.Table("player").Field("Nickname").Contains(nickname).Run(session)

	if err != nil {
		return false, err
	}

	var exists bool
	res.One(&exists)

	return exists, nil
}

func LoadPlayer(nickname string) (*Player, error) {
	res, err := r.Table("player").GetAllByIndex("Nickname", nickname).Run(session)

	if err != nil {
		return nil, err
	}

	if res.IsNil() {
		return nil, errors.New("Player does not exist")
	}

	var player Player
	res.One(&player)

	return &player, nil
}

func LoadPlayerByID(id string) (*Player, error) {
	res, err := r.Table("player").Get(id).Run(session)

	if err != nil {
		return nil, err
	}

	if res.IsNil() {
		return nil, errors.New("Player does not exist")
	}

	var player Player
	res.One(&player)

	return &player, nil
}

func CreatePlayer(nickname string) (*Player, error) {
	exists, err := NicknameExists(nickname)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("Nickname already exists.")
	}

	p := Player{Nickname: nickname, Created: time.Now()}

	_, err = r.Table("player").Insert(p).RunWrite(session)

	if err != nil {
		return nil, err
	}

	return LoadPlayer(nickname)
}

func (p *Player) Delete() error {
	_, err := r.Table("player").Get(p.Id).Delete().RunWrite(session)
	return err
}
