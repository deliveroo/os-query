package query

type noopMappable struct{}

func (n *noopMappable) Map() map[string]interface{} {
	return map[string]any{}
}

func NoopMappable() *noopMappable {
	return &noopMappable{}
}
