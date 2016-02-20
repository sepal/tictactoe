package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sepal/tictactoe/backend/game"
	"os"
)

var _ = Describe("Player", func() {
	var p1 *Player

	// Set up the database plus some demo data.
	BeforeEach(func() {
		var err error
		InitDB(os.Getenv("DB_HOST"), "tictactoe_test")

		p1, err = CreatePlayer("Player1")

		if err != nil {
			panic(err)
		}
	})

	// Remove the database.
	AfterEach(func() {
		err := p1.Delete()

		if err != nil {
			panic(err)
		}
	})

	Describe("A player", func() {
		It("should not be inserted if the nickname exists", func() {
			_, err := CreatePlayer("Player1")

			Expect(err).ToNot(BeNil())
		})
	})
})
