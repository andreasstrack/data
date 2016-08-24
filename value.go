package datastructures

// Key is an interface to represent
// the type of a key for a key/value store.
type Key interface {
	Comparable
	String() string
}

// Value is an interface to represent
// the type of a value for a key/value store.
type Value interface {
	String() string
}

// KeyValue is an interface for a key/value store.
type KeyValue interface {
	Key() Key
	Value() Value
}
