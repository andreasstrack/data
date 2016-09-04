package datastructures

import (
	"fmt"
	"reflect"
)

type Variant struct {
	val interface{}

	kind      reflect.Kind
	valueKind ValueKind

	b   bool
	i64 int64
	f64 float64
	s   string
}

func NewVariant(val interface{}) *Variant {
	v := new(Variant)
	v.Set(val)
	return v
}

func (v *Variant) Set(val interface{}) {
	vv := reflect.ValueOf(val)
	switch vv.Kind() {
	case reflect.Ptr:
		v.val = vv.Elem().Interface()
	default:
		v.val = vv.Interface()
	}
	v.kind = reflect.ValueOf(v.val).Kind()
	v.update()
}

func (v *Variant) Kind() reflect.Kind {
	return v.kind
}

func (v *Variant) ValueKind() ValueKind {
	return v.valueKind
}

func (v Variant) IsBool() bool {
	return v.valueKind == Bool
}

func (v Variant) Bool() bool {
	return v.b
}

func (v Variant) IsInt() bool {
	return v.valueKind == Int
}

func (v Variant) Int() int64 {
	return v.i64
}

func (v Variant) IsUint() bool {
	return v.valueKind == Uint
}

func (v Variant) Uint() uint64 {
	return uint64(v.i64)
}

func (v Variant) IsFloat() bool {
	return v.valueKind == Float
}

func (v Variant) Float() float64 {
	return v.f64
}

func (v Variant) IsString() bool {
	return v.valueKind == String
}

func (v Variant) String() string {
	switch v.valueKind {
	case Bool:
		return fmt.Sprintf("%s", v.b)
	case Int:
		return fmt.Sprintf("%d", v.i64)
	case Uint:
		return fmt.Sprintf("%d", (uint64)(v.i64))
	case Float:
		return fmt.Sprintf("%f", v.f64)
	case String:
		return v.s
	default:
		return fmt.Sprintf("%s", v.val)
	}
}

func (v Variant) LessThan(v2 Variant) bool {
	switch {
	case v.IsBool():
		return !v.Bool() && v2.Bool()
	case v.IsInt():
		return v.Int() < v2.Int()
	case v.IsUint():
		return v.Uint() < v2.Uint()
	case v.IsFloat():
		return v.Float() < v2.Float()
	case v.IsString():
		return v.String() < v2.String()
	default:
		panic("Invalid value kind for Variant")
	}
}

func (v Variant) EqualTo(v2 Variant) bool {
	switch {
	case v.IsBool():
		return v.Bool() == v2.Bool()
	case v.IsInt():
		return v.Int() == v2.Int()
	case v.IsUint():
		return v.Uint() == v2.Uint()
	case v.IsFloat():
		return v.Float() == v2.Float()
	case v.IsString():
		return v.String() == v2.String()
	default:
		panic("Invalid value kind for Variant")
	}
}

func (v *Variant) update() {
	switch v.kind {
	case reflect.Bool:
		v.b = v.val.(bool)
		v.valueKind = Bool
	case reflect.Int:
		v.i64 = int64(v.val.(int))
		v.valueKind = Int
	case reflect.Int8:
		v.i64 = int64(v.val.(int8))
		v.valueKind = Int
	case reflect.Int16:
		v.i64 = int64(v.val.(int16))
		v.valueKind = Int
	case reflect.Int32:
		v.i64 = int64(v.val.(int32))
		v.valueKind = Int
	case reflect.Int64:
		v.i64 = v.val.(int64)
		v.valueKind = Int
	case reflect.Uint:
		v.i64 = int64(v.val.(uint))
		v.valueKind = Uint
	case reflect.Uint8:
		v.i64 = int64(v.val.(uint8))
		v.valueKind = Uint
	case reflect.Uint16:
		v.i64 = int64(v.val.(uint16))
		v.valueKind = Uint
	case reflect.Uint32:
		v.i64 = int64(v.val.(uint32))
		v.valueKind = Uint
	case reflect.Uint64:
		v.i64 = int64(v.val.(uint64))
		v.valueKind = Uint
	case reflect.Float32:
		v.f64 = float64(v.val.(float32))
		v.valueKind = Float
	case reflect.Float64:
		v.f64 = v.val.(float64)
		v.valueKind = Float
	case reflect.String:
		v.s = v.val.(string)
		v.valueKind = String
	default:
		panic("Illegal value kind for Variant")
	}
}
