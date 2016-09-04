package datastructures

import "fmt"

type List interface {
	Add(e interface{})
	Remove(index int)
	At(index int) interface{}
	Front() interface{}
	Back() interface{}
	Clear()
	Size() int
	IsEmpty() bool
}

type ArrayList struct {
	list []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{list: make([]interface{}, 0)}
}

func (al *ArrayList) Add(e interface{}) {
	al.list = append(al.list, e)
}

func (al *ArrayList) Remove(index int) {
	al.list = append(al.list[:index], al.list[index+1:]...)
}

func (al *ArrayList) At(index int) interface{} {
	return al.list[index]
}

func (al *ArrayList) Front() interface{} {
	if len(al.list) == 0 {
		return nil
	}
	return al.list[0]
}

func (al *ArrayList) Back() interface{} {
	if len(al.list) == 0 {
		return nil
	}
	return al.list[len(al.list)-1]
}

func (al *ArrayList) Size() int {
	return len(al.list)
}

func (al *ArrayList) Clear() {
	al.list = make([]interface{}, 0)
}

func (al *ArrayList) IsEmpty() bool {
	return len(al.list) == 0
}

func (al *ArrayList) String() string {
	return fmt.Sprintf("%s", al.list)
}
