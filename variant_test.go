package datastructures

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	v := NewVariant(true)
	if v.Bool() != true {
		t.Errorf("%s != true", v)
		return
	}

	v.Set(false)
	if v.Bool() != false {
		t.Errorf("%s != false", v)
		return
	}
}

func TestInt(t *testing.T) {
	v := NewVariant(12)
	if v.Int() != 12 {
		t.Errorf("%s != 12", v)
		return
	}

	v.Set(int32(15))
	if v.Int() != 15 {
		t.Errorf("%s != 15", v)
		return
	}
}

func TestFloat(t *testing.T) {
	v := NewVariant(3.14)
	if v.Float() != 3.14 {
		t.Errorf("%s != 3.14", v)
		return
	}

	v.Set(-1000.1423875380275)
	if v.Float() != -1000.1423875380275 {
		t.Errorf("%s != -1000.1423875380275", v)
		return
	}
}

func TestString(t *testing.T) {
	v := NewVariant("Hello, world!")
	if v.String() != "Hello, world!" {
		t.Errorf("%s != 'Hello, world!", v)
		return
	}

	v.Set("Hello, You!")
	if v.String() != "Hello, You!" {
		t.Errorf("%s != 'Hello, You!", v)
		return
	}
}

func TestStringOutput(t *testing.T) {
	v := NewVariant(4)
	fmt.Printf("Variant with int(4): %s (%v)\n", v.String(), v)

	v = NewVariant("Hello")
	fmt.Printf("Variant with string(\"Hello\"): %s (%v)\n", v.String(), v)
}
