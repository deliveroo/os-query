package os_query

import "testing"

func TestRankFeature(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple ranking",
			RankFeature(
				"ranked_field", Log(2),
			).Boost(10),
			map[string]any{
				"rank_feature": map[string]any{
					"boost": 10,
					"field": "ranked_field",
					"log": map[string]any{
						"scaling_factor": 2,
					},
				},
			},
		},
	},
	)
}
