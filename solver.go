package bar_problem

type Color int

const (
	Unspecified Color = iota
	Red
	Blue
)

type Graph struct {
	verticesCount int
	adjacencyList [][]direction
}

type direction struct {
	to    int
	color Color
}

func NewGraph(vc int) *Graph {
	return &Graph{
		verticesCount: vc,
		adjacencyList: make([][]direction, vc),
	}
}

func (g *Graph) Connect(a, b int, color Color) {
	g.adjacencyList[a] = append(g.adjacencyList[a], direction{to: b, color: color})
	g.adjacencyList[b] = append(g.adjacencyList[b], direction{to: a, color: color})
}

func (g *Graph) Solve(from, to int) []int {
	visited := map[direction]bool{}
	return g.dfs(from, Unspecified, to, visited, []int{from})
}

func (g *Graph) getDirections(from int, curColor Color) []direction {
	var r []direction
	for _, dir := range g.adjacencyList[from] {

		switch dir.color {
		case Red:
			if curColor == Unspecified || curColor == Blue {
				r = append(r, dir)
			}
		case Blue:
			if curColor == Unspecified || curColor == Red {
				r = append(r, dir)
			}
		default:
			r = append(r, dir)
		}
	}

	return r
}

func (g *Graph) dfs(from int, curColor Color, dest int, visited map[direction]bool, curPath []int) []int {
	if from == dest {
		return curPath
	}

	curDir := direction{from, curColor}
	visited[curDir] = true
	for _, dir := range g.getDirections(from, curColor) {
		dir.color = applyColor(curColor, dir.color)
		if _, ok := visited[dir]; !ok {
			if _, ok := visited[dir]; !ok {
				if path := g.dfs(dir.to, dir.color, dest, visited, append(curPath, dir.to)); path != nil {
					return path
				}
			}
		}
	}
	delete(visited, curDir)

	return nil
}

func applyColor(curColor Color, targetColor Color) Color {
	if targetColor == Unspecified {
		return curColor
	}
	return targetColor
}
