package bar_problem

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

type connection struct {
	from int
	to int
	color Color
}

func TestGraph_Solve(t *testing.T) {
	testCases := []struct {
		gotVerticesCount int
		gotConnections []connection
		gotFrom int
		gotTo int
		want []int
	}{
		{
			gotVerticesCount: 20,
			gotConnections: []connection{
				{ from: 0, to: 1, color: Red },
				{ from: 1, to: 2, color: Red },
				{ from: 1, to: 3, color: Blue },
				{ from: 2, to: 4, color: Blue },
				{ from: 2, to: 5 },
				{ from: 3, to: 4, color: Blue },
				{ from: 3, to: 6, color: Red },
				{ from: 4, to: 5, color: Red },
				{ from: 4, to: 7, color: Red },
				{ from: 5, to: 8, color: Red },
				{ from: 6, to: 7, color: Blue },
				{ from: 7, to: 8, color: Blue },
				{ from: 7, to: 9, color: Blue },
				{ from: 8, to: 11 },
				{ from: 9, to: 10, color: Red },
				{ from: 10, to: 12, color: Blue },
				{ from: 11, to: 12, color: Red },
				{ from: 11, to: 14, color: Blue },
				{ from: 12, to: 13 },
				{ from: 12, to: 15, color: Red },
				{ from: 13, to: 16, color: Red },
				{ from: 14, to: 15, color: Red },
				{ from: 14, to: 17 },
				{ from: 15, to: 16, color: Blue },
				{ from: 15, to: 17 },
				{ from: 16, to: 18, color: Blue },
				{ from: 17, to: 18, color: Red },
				{ from: 18, to: 19, color: Red },
			},
			gotFrom: 0,
			gotTo: 19,
			want: []int{ 0, 1, 3, 6, 7, 4, 2, 5, 8, 11, 14, 15, 17, 14, 11, 8, 5, 2, 4, 7, 9, 10, 12, 13, 16, 18, 19},
		},
	}

	for _, testCase := range testCases {
		graph := NewGraph(testCase.gotVerticesCount)
		for _, connection := range testCase.gotConnections {
			graph.Connect(connection.from, connection.to, connection.color)
		}

		actual := graph.Solve(testCase.gotFrom, testCase.gotTo)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}