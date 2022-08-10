package user

import (
	"errors"
	"strings"
)

var ErrInvalidRoleName = errors.New("role has invalid role name")

type Role struct {
	RoleName string
}

func NewRole(roleName string) Role {
	return Role{RoleName: roleName}
}

func (r *Role) Validate() error {
	if strings.TrimSpace(r.RoleName) == "" {
		return ErrInvalidRoleName
	}

	return nil
}
