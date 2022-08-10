package pkg

import "testing"

func TestString(t *testing.T) {
	str, err := String(10)
	if err != nil {
		t.Fatalf("failed to create random string")
	}

	if len(str) != 10 {
		t.Fatalf("generated string length is wrong")
	}
}
