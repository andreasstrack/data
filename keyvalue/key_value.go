package keyvalue

import (
	"fmt"

	"github.com/andreasstrack/data"
)

// Key is an interface to represent
// the type of a key for a key/value store.
type Key interface {
	Int() int64
	String() string
}

// keyValue represents one element in a key-value-store.
type keyValue struct {
	k Key
	v data.Value
}

func (kv *keyValue) key() Key {
	return kv.k
}

func (kv *keyValue) value() data.Value {
	return kv.v
}

func newKeyValue(k Key, v data.Value) *keyValue {
	return &keyValue{k: k, v: v}
}

type KeyValueStore interface {
	Get(k Key) (data.Value, error)
	Set(k Key, v data.Value) error
	Contains(k Key) (bool, error)
	Delete(k Key) error
}

func Set(kvs KeyValueStore, k interface{}, v interface{}) error {
	return kvs.Set(data.NewVariant(k), data.NewVariant(v))
}

func Get(kvs KeyValueStore, k interface{}) (data.Value, error) {
	return kvs.Get(data.NewVariant(k))
}

func Delete(kvs KeyValueStore, k interface{}) error {
	return kvs.Delete(data.NewVariant(k))
}

type localKeyValueStoreWithIntKey struct {
	valueIndex map[int64]data.Value
}

func Contains(kvs KeyValueStore, k interface{}) (bool, error) {
	return kvs.Contains(data.NewVariant(k))
}

type localKeyValueStoreWithStringKey struct {
	valueIndex map[string]data.Value
}

func NewLocalKeyValueStore(kind data.ValueKind) KeyValueStore {
	switch kind {
	case data.Int:
		kvs := &localKeyValueStoreWithIntKey{}
		kvs.init()
		return kvs
	case data.String:
		kvs := &localKeyValueStoreWithStringKey{}
		kvs.init()
		return kvs
	}
	return nil
}

func (lkvi *localKeyValueStoreWithIntKey) init() {
	lkvi.valueIndex = make(map[int64]data.Value)
}

func (lkvi *localKeyValueStoreWithIntKey) Get(k Key) (data.Value, error) {
	var err error
	v, found := lkvi.valueIndex[k.Int()]
	if !found {
		err = fmt.Errorf("key '%d' not contained in store", k.Int())
	}
	return v, err
}

func (lkvi *localKeyValueStoreWithIntKey) Set(k Key, v data.Value) error {
	lkvi.valueIndex[k.Int()] = v
	return nil
}

func (lkvi *localKeyValueStoreWithIntKey) Contains(k Key) (bool, error) {
	_, found := lkvi.valueIndex[k.Int()]
	return found, nil
}

func (lkvi *localKeyValueStoreWithIntKey) Delete(k Key) error {
	delete(lkvi.valueIndex, k.Int())
	return nil
}

func (lkvs *localKeyValueStoreWithStringKey) init() {
	lkvs.valueIndex = make(map[string]data.Value)
}

func (lkvs *localKeyValueStoreWithStringKey) Get(k Key) (data.Value, error) {
	var err error
	v, found := lkvs.valueIndex[k.String()]
	if !found {
		err = fmt.Errorf("key '%s' not contained in store", k.String)
	}
	return v, err
}

func (lkvs *localKeyValueStoreWithStringKey) Set(k Key, v data.Value) error {
	lkvs.valueIndex[k.String()] = v
	return nil
}

func (lkvs *localKeyValueStoreWithStringKey) Delete(k Key) error {
	delete(lkvs.valueIndex, k.String())
	return nil
}

func (lkvs *localKeyValueStoreWithStringKey) Contains(k Key) (bool, error) {
	_, found := lkvs.valueIndex[k.String()]
	return found, nil
}
