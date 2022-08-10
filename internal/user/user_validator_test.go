package user

import "testing"

func TestNewUserValidator(t *testing.T) {
	minPassLen := uint(8)
	minUsernameLen := uint(10)
	v := NewUserValidator(minPassLen, minUsernameLen)

	t.Run("valid", func(t *testing.T) {
		usr := NewUser("username123", "password123")
		err := v.ValidateUser(usr)
		if err != nil {
			t.Fatalf("should be valid")
		}
	})

	t.Run("invalid", func(t *testing.T) {
		usr := NewUser("user", "pass")
		err := v.ValidateUser(usr)
		if err == nil {
			t.Fatalf("should be invalid")
		}
	})
}
