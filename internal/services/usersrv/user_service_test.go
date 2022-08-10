package usersrv

import (
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/store"
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
)

func TestUserService(t *testing.T) {
	roleStore := store.NewInMemoryRoleStore()
	roleSrv := rolesrv.NewRoleService(roleStore)
	usrStore := store.NewInMemoryUserStore()
	usrValidator := user.NewUserValidator(3, 3)
	srv := NewUserService(usrStore, roleSrv, usrValidator)

	username := "username"
	password := "password"
	roleName := "admin"

	_, err := roleSrv.CreateRole(roleName)
	if err != nil {
		t.Fatalf("failed to create role")
	}

	t.Run("create user", func(t *testing.T) {
		_, err := srv.CreateUser(username, password)
		if err != nil {
			t.Fatalf("failed to create user")
		}
	})

	t.Run("assign role", func(t *testing.T) {
		err := srv.AssignRole(username, roleName)
		if err != nil {
			t.Fatalf("failed to assign role")
		}
	})

	t.Run("check role", func(t *testing.T) {
		has, err := srv.UserHasRole(username, roleName)
		if err != nil {
			t.Fatalf("failed to check role")
		}
		if !has {
			t.Fatalf("role should be assigned")
		}
	})

	t.Run("get user", func(t *testing.T) {
		usr, err := srv.GetUserByUsername(username)
		if err != nil {
			t.Fatalf("failed to get user")
		}

		if usr.Username() != username || !usr.ComparePassword(password) {
			t.Fatalf("username and password should be equal")
		}
	})

	t.Run("delete user", func(t *testing.T) {
		err := srv.DeleteUser(username)
		if err != nil {
			t.Fatalf("failed to delete user")
		}

		_, err = srv.GetUserByUsername(username)
		if err == nil {
			t.Fatalf("user should be deleted")
		}
	})
}
