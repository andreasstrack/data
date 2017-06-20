package keyvalue

import (
	"fmt"
	"testing"
)

func SetGetTest(kvs KeyValueStore, k1, k2 interface{}, v interface{}, t *testing.T) {
	var err error

	vv := data.NewVariant(v)

	if err = Set(kvs, k1, v); err != nil {
		t.Error(err.Error())
	}

	if v2, err := Get(kvs, k1); err != nil {
		t.Error(err.Error())
	} else if v2.ValueKind() != vv.ValueKind() {
		t.Error(fmt.Errorf("ValueKind(%s) != %d"), v, vv.ValueKind())
	}

	if _, err = Get(kvs, k2); err == nil {
		t.Errorf("Expected error when reading key '%s' from %s", k2, kvs)
	}
}

func ContainsDeleteTest(kvs KeyValueStore, k interface{}, v interface{}, t *testing.T) {
	var err error

	if err = Set(kvs, k, v); err != nil {
		t.Error(err.Error())
		return
	}

	if contains, err := Contains(kvs, k); err != nil {
		t.Error(err.Error())
		return
	} else if !contains {
		t.Errorf("%s does not contain '%s'", kvs, k)
		return
	}

	if err = Delete(kvs, k); err != nil {
		t.Error(err.Error())
		return
	}

	if contains, err := Contains(kvs, k); err != nil {
		t.Error(err.Error())
		return
	} else if contains {
		t.Errorf("%s contain '%s' after delete", kvs, k)
		return
	}
}

func TestLocalWithStringKey(t *testing.T) {
	kvs := NewLocalKeyValueStore(data.String)
	SetGetTest(kvs, "Size", "Width", 5, t)
	ContainsDeleteTest(kvs, "Size", 5, t)
}

func TestLocalWithIntKey(t *testing.T) {
	kvs := NewLocalKeyValueStore(data.Int)
	SetGetTest(kvs, 5, 10, "Size", t)
	ContainsDeleteTest(kvs, 5, "Size", t)
}
