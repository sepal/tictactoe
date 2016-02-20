package game

import (
	r "github.com/dancannon/gorethink"
	"log"
)

var (
	session *r.Session
)

func InitDB(address, db string) {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  address,
		Database: db,
		MaxOpen:  40,
	})

	if err != nil {
		log.Fatalln(err.Error())
		panic(err)
	}

	res, err := r.DBList().Contains(db).Run(session)

	if err != nil {
		log.Fatalln(err.Error())
		panic(err)
	}

	var exists bool
	res.One(&exists)

	if exists == false {
		r.DBCreate(db).RunWrite(session)
		session.Use(db)
		CreateSchema()
	}

}

func DBClose()  {
	session.Close()
}

func CreateSchema() {
	tables := []string{
		"player",
		"player_session",
		"game",
	}

	for _, table := range tables {
		_, err := r.TableCreate(table).RunWrite(session)
		if err != nil {
			log.Fatalf("Error while creating table %v: %v", table, err.Error())
		}
	}
}
