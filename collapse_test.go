package query

import (
	"testing"
)

func TestCollapse_Map(t *testing.T) {
	runMapTests(
		t,
		[]mapTest{
			{
				name: "collapse",
				q:    Collapse("collapse_field"),
				exp: map[string]interface{}{
					"field": "collapse_field",
				},
			},
		},
	)
}
