package store

import "github.com/amanbolat/umsystem/internal/user"

type RoleDAO struct {
	RoleName string
}

func MapToRoleDAO(role user.Role) RoleDAO {
	return RoleDAO{RoleName: role.RoleName}
}

func MapToRole(dao RoleDAO) user.Role {
	return user.Role{RoleName: dao.RoleName}
}
