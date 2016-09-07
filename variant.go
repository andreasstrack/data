package datastructures

import (
	"fmt"
	"reflect"
)

type Variant struct {
	reflect.Value
	valueKind ValueKind
}

func NewVariant(val interface{}) *Variant {
	v := new(Variant)
	v.Set(val)
	return v
}

func (v *Variant) Set(val interface{}) {
	v.Value = reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v.Value = v.Elem()
	}
	v.update()
}

func (v *Variant) ValueKind() ValueKind {
	return v.valueKind
}

func (v *Variant) IsBool() bool {
	return v.valueKind == Bool
}

func (v *Variant) IsInt() bool {
	return v.valueKind == Int
}

func (v *Variant) IsUint() bool {
	return v.valueKind == Uint
}

func (v *Variant) IsFloat() bool {
	return v.valueKind == Float
}

func (v *Variant) IsString() bool {
	return v.valueKind == String
}

func (v *Variant) String() string {
	switch v.valueKind {
	case Bool:
		return fmt.Sprintf("%s", v.Bool())
	case Int:
		return fmt.Sprintf("%d", v.Int())
	case Uint:
		return fmt.Sprintf("%d", v.Uint())
	case Float:
		return fmt.Sprintf("%f", v.Float())
	case String:
		return v.Value.String()
	default:
		return v.Value.String()
	}
}

func (v *Variant) IsInterface() bool {
	return v.valueKind == Interface
}

func (v *Variant) LessThan(v2 Variant) bool {
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
	case v.IsInterface():
		panic("Cannot compare interface values (LessThan)")
	default:
		panic("Invalid value kind for Variant")
	}
}

func (v *Variant) EqualTo(v2 Variant) bool {
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
	case v.IsInterface():
		return v.Interface() == v2.Interface()
	default:
		panic("Invalid value kind for Variant")
	}
}

func (v *Variant) update() {
	switch v.Kind() {
	case reflect.Bool:
		v.valueKind = Bool
	case reflect.Int:
		v.valueKind = Int
	case reflect.Int8:
		v.valueKind = Int
	case reflect.Int16:
		v.valueKind = Int
	case reflect.Int32:
		v.valueKind = Int
	case reflect.Int64:
		v.valueKind = Int
	case reflect.Uint:
		v.valueKind = Uint
	case reflect.Uint8:
		v.valueKind = Uint
	case reflect.Uint16:
		v.valueKind = Uint
	case reflect.Uint32:
		v.valueKind = Uint
	case reflect.Uint64:
		v.valueKind = Uint
	case reflect.Float32:
		v.valueKind = Float
	case reflect.Float64:
		v.valueKind = Float
	case reflect.String:
		v.valueKind = String
	default:
		v.valueKind = Interface
	}
}
