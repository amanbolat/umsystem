package pkg

import "testing"

func TestCompareAndHashPassword(t *testing.T) {
	password := "password"

	hash := HashPassword(password)

	if !ComparePassword(hash, password) {
		t.Fatalf("passwords should be equal")
	}

	if ComparePassword(hash, "wrong") {
		t.Fatalf("compare should fail")
	}
}
