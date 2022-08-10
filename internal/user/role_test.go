package user

import "testing"

func TestRole_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		role := NewRole("admin")
		err := role.Validate()
		if err != nil {
			t.Fatalf("should be valid")
		}
	})

	t.Run("invalid", func(t *testing.T) {
		role := NewRole("   ")
		err := role.Validate()
		if err == nil {
			t.Fatalf("should be invalid")
		}
	})
}
