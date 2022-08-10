package store

import (
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
)

func TestInMemoryRoleStore(t *testing.T) {
	store := NewInMemoryRoleStore()
	role := user.NewRole("admin")

	t.Run("insert role", func(t *testing.T) {
		err := store.InsertRole(role)
		if err != nil {
			t.Fatalf("failed to insert role")
		}
	})

	t.Run("get role", func(t *testing.T) {
		_, err := store.GetRoleByName(role.RoleName)
		if err != nil {
			t.Fatalf("failed to find role")
		}
	})

	t.Run("delete role", func(t *testing.T) {
		err := store.DeleteRole(role.RoleName)
		if err != nil {
			t.Fatalf("faield to delete role")
		}

		_, err = store.GetRoleByName(role.RoleName)
		if err == nil {
			t.Fatalf("role was not deleted")
		}
	})
}
