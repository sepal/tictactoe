package game_test

import (
	. "github.com/sepal/tictactoe/backend/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var p1, p2 *Player
	var g *Game

	BeforeEach(func() {
		p1, _ = CreatePlayer("Player1")
		p2, _ = CreatePlayer("Player2")

		g = CreateGame(p1, p2)
	})

	AfterEach(func() {
		p1.Delete()
		p2.Delete()
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
