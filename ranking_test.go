package esquery

import "testing"

func TestRankFeature(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple ranking",
			RankFeature(
				"ranked_filed", Log(2),
			),
			map[string]any{
				"boost": 1,
				"field": "ranked_filed",
				"log": map[string]any{
					"scaling_factor": 2,
				},
			},
		},
	},
	)
}
