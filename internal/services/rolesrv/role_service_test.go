package rolesrv

import (
	"github.com/amanbolat/umsystem/internal/store"
	"testing"
)

func TestRoleService(t *testing.T) {
	roleStore := store.NewInMemoryRoleStore()
	srv := NewRoleService(roleStore)

	roleName := "admin"

	t.Run("create role", func(t *testing.T) {
		role, err := srv.CreateRole(roleName)
		if err != nil {
			t.Fatalf("failed to create role")
		}
		if role.RoleName != roleName {
			t.Fatalf("wrong role name")
		}
	})

	t.Run("get role", func(t *testing.T) {
		role, err := srv.GetRoleByName(roleName)
		if err != nil {
			t.Fatalf("failed to get role")
		}
		if role.RoleName != roleName {
			t.Fatalf("wrong role name")
		}
	})

	t.Run("delete role", func(t *testing.T) {
		err := srv.DeleteRole(roleName)
		if err != nil {
			t.Fatalf("failed to delete role")
		}
		_, err = srv.GetRoleByName(roleName)
		if err == nil {
			t.Fatalf("role should be deleted")
		}
	})
}
