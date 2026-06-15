package main

import "testing"

func TestAdd(t *testing.T) {

	m := NewStringIntMap()
	m.Add("key", 42)

	val, ok := m.Get("key")
	if !ok {
		t.Errorf("Add() - ключ не найден после добавления")
	}
	if val != 42 {
		t.Errorf("Add() = %d, want 42", val)
	}
}

func TestRemove(t *testing.T) {

	m := NewStringIntMap()
	m.Add("key", 42)
	m.Add("key2", 43)
	m.Remove("key")

	if m.Exists("key") {
		t.Errorf("Remove() - ключ должен быть удалён")
	}

	val, ok := m.Get("key2")
	if !ok {
		t.Errorf("Remove() - второй ключ пропал")
	}
	if val != 43 {
		t.Errorf("Remove() = %d, want 43", val)
	}
}

func TestCopy(t *testing.T) {

	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	copyMap := m.Copy()

	if len(copyMap) != 2 {
		t.Errorf("TestCopy() lenght = %d,want 2", len(copyMap))
	}

	if copyMap["a"] != 1 {
		t.Errorf("TestCopy()[a] = %d, want 1", copyMap["a"])
	}

	if copyMap["b"] != 2 {
		t.Errorf("TestCopy()[a] = %d, want 2", copyMap["b"])
	}
	copyMap["a"] = 999
	val, _ := m.Get("a")
	if val != 1 {
		t.Errorf("Copy() - копия влияет на оригинал")
	}
}

func TestExists(t *testing.T) {

	m := NewStringIntMap()
	m.Add("a", 1)

	if !m.Exists("a") {
		t.Errorf("Exist() - ключ должен существовать")
	}

	if m.Exists("b") {
		t.Errorf("Exist() - ключа не должно быть")
	}
}

func TestGet(t *testing.T) {

	m := NewStringIntMap()
	m.Add("key", 1)

	val, ok := m.Get("key")
	if !ok {
		t.Errorf("Get() - ключ не найден")
	}
	if val != 1 {
		t.Errorf("Get() =%d, want 1", val)
	}

	_, ok = m.Get("key123")
	if ok {
		t.Errorf("Get() - такого ключа быть не должно, ok должно быть false")
	}
}
