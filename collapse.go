package query

// CollapseRequest represents the "collapse" param that can be applied to a request
// see: https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html
type CollapseRequest struct {
	field string
}

// Collapse creates a new collapse request
func Collapse(field string) *CollapseRequest {
	return &CollapseRequest{
		field: field,
	}
}

// Map returns a map representation of the Source object.
func (source CollapseRequest) Map() map[string]interface{} {
	return map[string]interface{}{
		"field": source.field,
	}
}
