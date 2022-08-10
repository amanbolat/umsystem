package store

import (
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
)

func TestMapDAOToUser(t *testing.T) {
	dao := RoleDAO{RoleName: "admin"}
	role := MapToRole(dao)
	if dao.RoleName != role.RoleName {
		t.Fatalf("dao and role rolenames are not equal")
	}
}

func TestMapUserToDAO(t *testing.T) {
	usr := user.NewUser("user", "pass")
	dao := MapUserToDAO(usr)
	if dao.Password != usr.Password() {
		t.Fatalf("dao and user passwords are not equal")
	}

	if dao.Username != usr.Username() {
		t.Fatalf("dao and user passwords are not equal")
	}

	for _, role := range dao.Roles {
		if !usr.HasRole(role) {
			t.Fatalf("user should have the role: %s", role)
		}
	}
}
