package data

import "reflect"

// Value is an interface to represent
// the type of a value for a key/value store.
type Value interface {
	//	IsBool() bool
	Bool() bool
	IsInt() bool
	Int() int64
	IsUint() bool
	Uint() uint64
	IsFloat() bool
	Float() float64
	IsString() bool
	String() string
	Interface() interface{}
	ReflectValue() *reflect.Value
}

func IsSimpleData(v Value) bool {
	switch {
	//	case v.IsInt():
	//		return true
	case v.IsFloat():
		return true
	case v.IsString():
		return true
	default:
		return false
	}
}

// A ComparableValue is composed of a Comparable and a Value.
type ComparableValue interface {
	Comparable
	Value
}
