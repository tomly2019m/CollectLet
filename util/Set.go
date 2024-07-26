package util

type Set struct {
	elements map[interface{}]interface{}
}

func NewSet() *Set {
	return &Set{make(map[interface{}]interface{})}
}

func (set *Set) Add(element interface{}) {
	set.elements[element] = nil
}

func (set *Set) Remove(element interface{}) {
	delete(set.elements, element)
}

func (set *Set) Contains(element interface{}) bool {
	_, ok := set.elements[element]
	return ok
}

func (set *Set) Size() int {
	return len(set.elements)
}

func (set *Set) Iterate(f func(interface{})) {
	for element := range set.elements {
		f(element)
	}
}
