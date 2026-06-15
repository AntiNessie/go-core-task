package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {

	s.data[key] = value

}

func (s *StringIntMap) Remove(key string) {

	delete(s.data, key)

}

func (s *StringIntMap) Copy() map[string]int {

	result := make(map[string]int)

	for k, v := range s.data {
		result[k] = v
	}

	return result

}

func (s *StringIntMap) Exists(key string) bool {

	_, ok := s.data[key]
	return ok

}

func (s *StringIntMap) Get(key string) (int, bool) {

	v, ok := s.data[key]
	return v, ok

}

func main() {
	m := NewStringIntMap()

	m.Add("one", 1)
	m.Add("two", 2)

	val, ok := m.Get("one")
	fmt.Printf("Get one: %d, %v\n", val, ok)

	fmt.Printf("Exists two: %v\n", m.Exists("two"))

	copied := m.Copy()
	fmt.Printf("Copied: %v\n", copied)

	m.Remove("one")
	fmt.Printf("Exists one after remove: %v\n", m.Exists("one"))
}
