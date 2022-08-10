package store

import "github.com/amanbolat/umsystem/internal/user"

type UserDAO struct {
	Username string
	Password string
	Roles    []string
}

func MapDAOToUser(dao UserDAO) user.User {
	usr := user.NewUser(dao.Username, dao.Password)
	for _, role := range dao.Roles {
		usr.AssignRole(role)
	}
	return usr
}

func MapUserToDAO(usr user.User) UserDAO {
	roles := make([]string, len(usr.Roles()))
	copy(roles, usr.Roles())

	return UserDAO{
		Username: usr.Username(),
		Password: usr.Password(),
		Roles:    roles,
	}
}
