package datastructures

import "reflect"

// Value is an interface to represent
// the type of a value for a key/value store.
type Value interface {
	Bool() bool
	Int() int64
	Uint() uint64
	Float() float64
	String() string
	Interface() interface{}
	ReflectValue() *reflect.Value
}

// A ComparableValue is composed of a Comparable and a Value.
type ComparableValue interface {
	Comparable
	Value
}
