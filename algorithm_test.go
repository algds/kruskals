package kruskals

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	graph, mst []WeightedEdge
}{
	{
		[]WeightedEdge{
			SimpleWeightedEdge{0, 3, 3},
			SimpleWeightedEdge{3, 1, 30},
			SimpleWeightedEdge{0, 1, 20},
			SimpleWeightedEdge{0, 4, 10},
			SimpleWeightedEdge{1, 4, 5},
			SimpleWeightedEdge{4, 2, 20},
			SimpleWeightedEdge{1, 2, 50},
		},
		[]WeightedEdge{
			SimpleWeightedEdge{0, 3, 3},
			SimpleWeightedEdge{1, 4, 5},
			SimpleWeightedEdge{0, 4, 10},
			SimpleWeightedEdge{4, 2, 20},
		},
	},
}

func TestMaximumSpanningTree(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinimumSpanningTree(tc.graph); !reflect.DeepEqual(result, tc.mst) {
			t.Errorf("Expected %v, got %v", tc.mst, result)
		}
	}
}

func BenchmarkMaximumSpanningTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumSpanningTree(tc.graph)
		}
	}
}
