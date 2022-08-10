package user

import "github.com/amanbolat/umsystem/pkg"

type User struct {
	username string
	password string
	roles    []string
}

func (u *User) EqualTo(u2 User) bool {
	if u.password != u2.password {
		return false
	}

	if u.username != u2.username {
		return false
	}

	if len(u.roles) != len(u2.roles) {
		return false
	}

	for _, r := range u.roles {
		if !u2.HasRole(r) {
			return false
		}
	}

	return true
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Roles() []string {
	return u.roles
}

func (u *User) HasRole(role string) bool {
	for _, r := range u.roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) AssignRole(role string) {
	for _, r := range u.roles {
		if role == r {
			return
		}
	}
	u.roles = append(u.roles, role)
}

func (u *User) PreSave() {
	u.password = pkg.HashPassword(u.password)
}

func (u *User) ComparePassword(password string) bool {
	return pkg.ComparePassword(u.password, password)
}

func NewUser(username string, password string) User {
	return User{
		username: username,
		password: password,
	}
}
