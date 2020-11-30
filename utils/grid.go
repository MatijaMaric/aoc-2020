package utils

import "math"

// Pos 2D vector
type Pos struct {
	X, Y int
}

// ManhattanDist return Manhattan distance of vectors
func (p Pos) ManhattanDist(q Pos) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

// EuclidDist returns Euclidian distance of vectors
func (p Pos) EuclidDist(q Pos) float64 {
	return math.Hypot(float64(Abs(p.X-q.X)), float64(Abs(p.Y-q.Y)))
}

// Add adds vectors
func (p Pos) Add(q Pos) Pos {
	return Pos{p.X + q.X, p.Y + q.Y}
}

// Moves returns single step moves in all directions on a grid
func (p Pos) Moves() []Pos {
	return []Pos{
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
	}
}

// BoundingBox returns boundaries (min and max) of vectors
func BoundingBox(ps []Pos) (min, max Pos) {
	min.X, min.Y = MaxInt, MaxInt
	max.X, max.Y = MinInt, MinInt
	for _, p := range ps {
		min.X = Min(min.X, p.X)
		min.Y = Min(min.Y, p.Y)
		max.X = Max(max.X, p.X)
		max.Y = Max(max.Y, p.Y)
	}
	return
}

// Grid grid
type Grid struct {
	width, height int
}

// Valid returns true if vector is inside grid
func (g Grid) Valid(p Pos) bool {
	return p.X >= 0 && p.X < g.width && p.Y >= 0 && p.Y < g.height
}

// Maze is a grid with walls
type Maze struct {
	Grid
	IsWall func(Pos) bool
}

// Valid returns if vector is walkable
func (m Maze) Valid(p Pos) bool {
	return m.Grid.Valid(p) && !m.IsWall(p)
}

// ValidMoves returns all valid moves in maze
func (p Pos) ValidMoves(m Maze) []Pos {
	var valid []Pos
	for _, move := range p.Moves() {
		if m.Valid(move) {
			valid = append(valid, move)
		}
	}
	return valid
}

// ShortestPath return shortest path from-to using BFS
func (m Maze) ShortestPath(from, to Pos) []Pos {
	type entry struct {
		Pos
		dist int
		path []Pos
	}

	var queue []entry
	queue = append(queue, entry{from, 0, []Pos{}})

	seen := make(map[Pos]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[:len(queue)-1]

		if _, ok := seen[current.Pos]; ok {
			continue
		}

		if current.Pos == to {
			return current.path
		}

		seen[current.Pos] = true

		for _, next := range current.Pos.ValidMoves(m) {
			queue = append(queue, entry{next, current.dist + 1, append(current.path, next)})
		}
	}

	return []Pos{}
}
