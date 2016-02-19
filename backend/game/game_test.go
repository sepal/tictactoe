package game

import (
	"testing"
	. "github.com/franela/goblin"
	"fmt"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Game", func() {
		g.It("Should not allow same move", func() {
			p1 := Player{0, "player1"}
			p2 := Player{1, "player2"}

			game := CreateGame(&p1, &p2)

			err := game.TakeTurn(Vertex{0, 0})
			fmt.Println(game.actions)
			g.Assert(err == nil).IsTrue("There was an error on the first move!")

			err = game.TakeTurn(Vertex{0, 0})
			g.Assert(err != nil).IsTrue("There was no error, although its the same move!")
		})

		g.It("Should win on column", func() {
			p1 := Player{0, "player1"}
			p2 := Player{1, "player2"}

			game := CreateGame(&p1, &p2)

			g.Assert(game.state == STATE_RUNNING).IsTrue("Game is not running!")

			// p1 turn
			game.TakeTurn(Vertex{0, 0});
			// p2 turn
			game.TakeTurn(Vertex{2, 0});
			// p1 turn
			game.TakeTurn(Vertex{0, 1});
			// p2 turn
			game.TakeTurn(Vertex{2, 1});
			// p1 turn
			game.TakeTurn(Vertex{0, 2});

			g.Assert(game.state == STATE_FINISHED).IsTrue("Game not finished!")
			g.Assert(game.current == &p1).IsTrue("Player 1 has not won!")
		})

		g.It("Should win on row", func() {
			p1 := Player{0, "player1"}
			p2 := Player{1, "player2"}

			game := CreateGame(&p1, &p2)

			g.Assert(game.state == STATE_RUNNING).IsTrue("Game is not running!")

			// p1 turn
			g.Assert(game.current == &p1).IsTrue("Its not player1's turn!")
			game.TakeTurn(Vertex{0, 0});
			// p2 turn
			g.Assert(game.current == &p2).IsTrue("Its not player2's turn!")
			game.TakeTurn(Vertex{1, 0});
			// p1 turn
			g.Assert(game.current == &p1).IsTrue("Its not player1's turn!")
			game.TakeTurn(Vertex{0, 1});
			// p2 turn
			g.Assert(game.current == &p2).IsTrue("Its not player2's turn!")
			game.TakeTurn(Vertex{1, 1});
			// p1 turn
			g.Assert(game.current == &p1).IsTrue("Its not player1's turn!")
			game.TakeTurn(Vertex{2, 0});
			// p2 turn
			g.Assert(game.current == &p2).IsTrue("Its not player2's turn!")
			game.TakeTurn(Vertex{1, 2})


			g.Assert(game.state == STATE_FINISHED).IsTrue("Game not finished!")
			g.Assert(game.current == &p2).IsTrue("Player 2 has not won!" + game.current.nickname)
		})

		g.It("Should win on diagonal", func() {
			p1 := Player{0, "player1"}
			p2 := Player{1, "player2"}

			game := CreateGame(&p1, &p2)

			g.Assert(game.state == STATE_RUNNING).IsTrue("Game is not running!")

			// p1 turn
			game.TakeTurn(Vertex{1, 1});
			// p2 turn
			game.TakeTurn(Vertex{2, 1});
			// p1 turn
			game.TakeTurn(Vertex{0, 0});
			// p2 turn
			game.TakeTurn(Vertex{2, 0});
			// p1 turn
			game.TakeTurn(Vertex{2, 2});

			g.Assert(game.state == STATE_FINISHED).IsTrue("Game not finished!")
			g.Assert(game.current == &p1).IsTrue("Player 1 has not won!")
		})
	})
}
