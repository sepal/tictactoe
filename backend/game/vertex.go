package game

type Vertex struct {
	X, Y int
}

func (v *Vertex) Equals(v2 Vertex) bool {
	return v.X == v2.X && v.Y == v2.Y
}