package user

import "testing"

func TestUser(t *testing.T) {
	t.Run("new user", func(t *testing.T) {
		username := "user"
		password := "password"
		usr := NewUser(username, password)
		if usr.Password() != password {
			t.Fatalf("passwords are not equal")
		}

		if usr.Username() != username {
			t.Fatalf("usernames are not equal")
		}
	})

	t.Run("assign role and check", func(t *testing.T) {
		usr := NewUser("user", "password")
		role := "admin"
		usr.AssignRole(role)
		if !usr.HasRole(role) {
			t.Fatalf("failed to assign role")
		}
	})

	t.Run("pre save and compare password hash", func(t *testing.T) {
		username := "user"
		password := "password"
		usr := NewUser(username, password)
		usr.PreSave()
		if password == usr.Password() {
			t.Fatalf("user presave failed")
		}

		if !usr.ComparePassword(password) {
			t.Fatalf("passwords are not equal")
		}
	})
}
