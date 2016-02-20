package game_test

import (
	. "github.com/sepal/tictactoe/backend/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Game", func() {
	var p1, p2 *Player
	var g *Game

	BeforeEach(func() {
		var err error

		InitDB(os.Getenv("DB_HOST"), "tictactoe_test")

		p1, err = CreatePlayer("Player1")

		if err != nil {
			panic(err)
		}

		p2, err = CreatePlayer("Player2")

		if err != nil {
			panic(err)
		}

		g = CreateGame(p1, p2)
	})

	AfterEach(func() {
		p1.Delete()
		p2.Delete()

		DBClose()
	})

	Describe("A new game", func() {
		It("should not allow the same move", func() {
			err := g.TakeTurn(Vertex{0, 0})
			Expect(err).To(BeNil())
		})

		It("should win on column", func() {
			Expect(g.GetState()).To(Equal(STATE_RUNNING))

			// p1 turn
			g.TakeTurn(Vertex{0, 0})
			// p2 turn
			g.TakeTurn(Vertex{2, 0})
			// p1 turn
			g.TakeTurn(Vertex{0, 1})
			// p2 turn
			g.TakeTurn(Vertex{2, 1})
			// p1 turn
			g.TakeTurn(Vertex{0, 2})

			Expect(g.GetState()).To(Equal(STATE_FINISHED))
			Expect(g.GetCurrentPlayer()).To(Equal(p1))
		})

		It("Should win on row", func() {
			Expect(g.GetState()).To(Equal(STATE_RUNNING))

			// p1 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p1))
			g.TakeTurn(Vertex{0, 0})
			// p2 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p2))
			g.TakeTurn(Vertex{1, 0})
			// p1 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p1))
			g.TakeTurn(Vertex{0, 1})
			// p2 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p2))
			g.TakeTurn(Vertex{1, 1})
			// p1 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p1))
			g.TakeTurn(Vertex{2, 0})
			// p2 turn
			Expect(g.GetCurrentPlayer()).To(Equal(p2))
			g.TakeTurn(Vertex{1, 2})

			Expect(g.GetState()).To(Equal(STATE_FINISHED))
			Expect(g.GetCurrentPlayer()).To(Equal(p2))
		})

		It("should win on diagonal", func() {
			Expect(g.GetState()).To(Equal(STATE_RUNNING))

			// P1 turn
			g.TakeTurn(Vertex{1, 1})
			// P2 turn
			g.TakeTurn(Vertex{2, 1})
			// P1 turn
			g.TakeTurn(Vertex{0, 0})
			// P2 turn
			g.TakeTurn(Vertex{2, 0})
			// P1 turn
			g.TakeTurn(Vertex{2, 2})

			Expect(g.GetState()).To(Equal(STATE_FINISHED))
			Expect(g.GetCurrentPlayer()).To(Equal(p1))
		})
	})
})
