package query

// Collapse represents the "collapse" param that can be applied to a request
// see: https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html
type Collapse struct {
	field string
}

// Map returns a map representation of the Source object.
func (source Collapse) Map() map[string]interface{} {
	return map[string]interface{}{
		"field": source.field,
	}
}
