package os_query

// Mappable is the interface implemented by the various query and aggregation
// types provided by the package. It allows the library to easily transform the
// different queries to "generic" maps that can be easily encoded to JSON.
type Mappable interface {
	Map() map[string]interface{}
}

// Aggregation is an interface that each aggregation type must implement. It
// is simply an extension of the Mappable interface to include a Named function,
// which returns the name of the aggregation.
type Aggregation interface {
	Mappable
	Name() string
}
