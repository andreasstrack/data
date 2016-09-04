package datastructures

type ValueKind int64

const (
	Bool ValueKind = 1 << iota
	Int
	Uint
	Float
	String
)

// Value is an interface to represent
// the type of a value for a key/value store.
type Value interface {
	ValueKind() ValueKind
	Bool() bool
	Int() int64
	Uint() uint64
	Float() float64
	String() string
}

// A ComparableValue is composed of a Comparable and a Value.
type ComparableValue interface {
	Comparable
	Value
}
