package main

import (
	"testing"
)

func TestGetNeighbours(t *testing.T) {
	var neighbourTests = []struct {
		name       string
		target     *Node
		neighbours []*Node
		expected   []*Node
	}{
		{"testing obstacle tile nodes",
			&Node{X: 3, Y: 3},
			[]*Node{&Node{}, &Node{}, &Node{}, &Node{Type: Obstacle}},
			[]*Node{&Node{}, &Node{}, &Node{}},
		},

		{"testing out of bounds neighbours",
			&Node{X: 0, Y: 0},
			[]*Node{&Node{}, &Node{}, &Node{}, &Node{}},
			[]*Node{&Node{}, &Node{}},
		},
	}
	for _, tc := range neighbourTests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewMap(5, 5)
			x, y := tc.target.X, tc.target.Y
			m[x][y] = tc.target
			if x-1 >= 0 {
				m[x-1][y] = tc.neighbours[0]
			}
			if y-1 >= 0 {
				m[x][y-1] = tc.neighbours[1]
			}
			if x+1 < len(m) {
				m[x+1][y] = tc.neighbours[2]
			}
			if y+1 < len(m[0]) {
				m[x][y+1] = tc.neighbours[3]
			}

			nbs := m.GetNeighbours(tc.target)
			if !SliceEq(t, nbs, tc.expected) {
				t.Errorf("Testcase: %q: the wrong list of neighbours was returned %s \n instead of %s", tc.name, nbs, tc.expected)
			}
		})
	}
}

func SliceEq(t *testing.T, a, b []*Node) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if *(a[i]) != *(b[i]) {
			return false
		}
	}

	return true
}
