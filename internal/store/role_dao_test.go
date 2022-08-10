package store

import (
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
)

func TestMapToRole(t *testing.T) {
	dao := RoleDAO{RoleName: "admin"}
	role := MapToRole(dao)
	if dao.RoleName != role.RoleName {
		t.Fatalf("dao and role rolenames are not equal")
	}
}

func TestMapToRoleDAO(t *testing.T) {
	role := user.NewRole("admin")
	dao := MapToRoleDAO(role)
	if dao.RoleName != role.RoleName {
		t.Fatalf("dao and role rolenames are not equal")
	}
}
