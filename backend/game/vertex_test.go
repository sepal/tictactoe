package game

import "testing"

func TestEquals(t *testing.T) {
	v1 := Vertex{1, 0}
	v2 := Vertex{0, 1}

	if v1.Equals(v2) {
		t.Fatalf("Vertex %v,%v equals vertex %v,%v", v1.X, v1.Y, v2.X, v2.Y)
	}

	v1 = Vertex{5, 10}
	v2 = Vertex{5, 10}

	if !v1.Equals(v2) {
		t.Fatalf("Vertex %v,%v not equals vertex %v,%v", v1.X, v1.Y, v2.X, v2.Y)
	}
}
