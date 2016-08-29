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
	String() string
}

type BoolValue interface {
	Value
	Bool()
}

type IntValue interface {
	Value
	Int() int64
}

type UintValue interface {
	Value
	Uint() uint64
}

type FloatValue interface {
	Value
	Float() float64
}

type StringValue interface {
	Value
}

// A ComparableValue is composed of a Comparable and a Value.
type ComparableValue interface {
	Comparable
	Value
}
