package rolesrv

import "github.com/amanbolat/umsystem/internal/user"

type RoleStore interface {
	InsertRole(role user.Role) error
	DeleteRole(roleName string) error
	GetRoleByName(roleName string) (user.Role, error)
}

type RoleService interface {
	CreateRole(roleName string) (user.Role, error)
	DeleteRole(roleName string) error
	GetRoleByName(roleName string) (user.Role, error)
}

type roleService struct {
	store RoleStore
}

func NewRoleService(store RoleStore) RoleService {
	return &roleService{store: store}
}

func (r *roleService) GetRoleByName(roleName string) (user.Role, error) {
	return r.store.GetRoleByName(roleName)
}

func (r *roleService) CreateRole(roleName string) (user.Role, error) {
	role := user.NewRole(roleName)
	err := role.Validate()
	if err != nil {
		return user.Role{}, err
	}

	err = r.store.InsertRole(role)
	if err != nil {
		return user.Role{}, err
	}
	return role, nil
}

func (r *roleService) DeleteRole(roleName string) error {
	return r.store.DeleteRole(roleName)
}
