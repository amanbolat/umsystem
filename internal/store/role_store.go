package store

import (
	"errors"
	"github.com/amanbolat/umsystem/internal/user"
	"github.com/amanbolat/umsystem/pkg"
)

var ErrRoleAlreadyExists = errors.New("role does already exist")
var ErrRoleNotFound = errors.New("role was not found")

type InMemoryRoleStore struct {
	m *pkg.Map[RoleDAO]
}

func NewInMemoryRoleStore() *InMemoryRoleStore {
	return &InMemoryRoleStore{m: pkg.NewMap[RoleDAO]()}
}

func (s *InMemoryRoleStore) InsertRole(role user.Role) error {
	if s.m.Exists(role.RoleName) {
		return ErrRoleAlreadyExists
	}

	s.m.Set(role.RoleName, MapToRoleDAO(role))
	return nil
}

func (s *InMemoryRoleStore) DeleteRole(roleName string) error {
	if !s.m.Exists(roleName) {
		return ErrRoleNotFound
	}

	s.m.Delete(roleName)
	return nil
}

func (s *InMemoryRoleStore) GetRoleByName(roleName string) (user.Role, error) {
	dao, ok := s.m.Get(roleName)
	if !ok {
		return user.Role{}, ErrRoleNotFound
	}

	return MapToRole(dao), nil
}
