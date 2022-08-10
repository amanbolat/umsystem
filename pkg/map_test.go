package pkg

import (
	"testing"
)

type T struct {
	a int
	b int
}

func (t T) Equal(t2 T) bool {
	return t.a == t2.a && t.b == t2.b
}

func TestMap(t *testing.T) {
	m := NewMap[*T]()

	one := &T{a: 1, b: 2}

	m.Set("one", one)

	t.Run("get non-existing value", func(t *testing.T) {
		val, ok := m.Get("two")
		if val != nil || ok {
			t.Fatal("value for key two should not exist")
		}
	})

	t.Run("get existing value", func(t *testing.T) {
		val, ok := m.Get("one")
		if val == nil || !ok {
			t.Fatalf("value for key one is nil")
		}

		if !val.Equal(*one) {
			t.Fatalf("retrieved value is not equal to the actual one: [%v] != [%v]", *val, *one)
		}
	})

	t.Run("get existing value", func(t *testing.T) {
		ok := m.Exists("one")
		if !ok {
			t.Fatal("value one should exist")
		}
	})

	t.Run("delete value", func(t *testing.T) {
		m.Delete("one")
		ok := m.Exists("one")
		if ok {
			t.Fatal("value one should not exist as it was deleted")
		}
	})
}
