package store

import (
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
)

func TestInMemoryUserStore(t *testing.T) {
	store := NewInMemoryUserStore()

	user1 := user.NewUser("user", "pass")
	user1.AssignRole("admin")

	t.Run("insert user", func(t *testing.T) {
		err := store.InsertUser(user1)
		if err != nil {
			t.Fatalf("failed to insert user")
		}
	})

	t.Run("get user", func(t *testing.T) {
		usr, err := store.GetUserByUsername(user1.Username())
		if err != nil {
			t.Fatalf("failed to find user")
		}
		if !usr.EqualTo(user1) {
			t.Fatalf("fonud user is not equal to the original one")
		}
	})

	t.Run("update user", func(t *testing.T) {
		updated := user.NewUser("user", "new_pass")
		err := store.UpdateUser(updated)
		if err != nil {
			t.Fatalf("failed to update user")
		}

		found, _ := store.GetUserByUsername(updated.Username())
		if !found.EqualTo(updated) {
			t.Fatalf("failed to udpate user")
		}
	})

	t.Run("delete user", func(t *testing.T) {
		err := store.DeleteUser(user1.Username())
		if err != nil {
			t.Fatalf("failed to delete user")
		}

		_, err = store.GetUserByUsername(user1.Username())
		if err == nil {
			t.Fatalf("user should be deleted")
		}
	})
}
