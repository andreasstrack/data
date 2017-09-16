package set

import (
	R "github.com/andreasstrack/util/reflect"
)

type Set struct {
	content map[interface{}]bool
}

func NewSet() (*Set, error) {
	s := &Set{}
	s.content = make(map[interface{}]bool)
	return s, nil
}

func NewSetFromSlice(slice interface{}) (*Set, error) {
	var err error
	var sliceOfInterface *[]interface{}
	var s *Set

	if sliceOfInterface, err = R.GetSlice(slice); err != nil {
		return nil, err
	}
	if s, err = NewSet(); err != nil {
		return nil, err
	}

	for i := range *sliceOfInterface {
		s.Insert((*sliceOfInterface)[i])
	}
	return s, nil
}

func (s *Set) Insert(v interface{}) {
	s.content[v] = true
}

func (s *Set) Remove(v interface{}) {
	delete(s.content, v)
}

func (s Set) Contains(v interface{}) bool {
	_, found := s.content[v]
	return found
}

func (s Set) Size() int {
	return len(s.content)
}

func (s Set) IsEmpty() bool {
	return s.Size() == 0
}
