package store

import (
	"errors"
	"github.com/amanbolat/umsystem/internal/user"
	"github.com/amanbolat/umsystem/pkg"
)

var ErrUserNotFound = errors.New("user was not found")
var ErrUserAlreadyExists = errors.New("user with given username does already exist")

type InMemoryUserStore struct {
	m *pkg.Map[UserDAO]
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{m: pkg.NewMap[UserDAO]()}
}

func (i *InMemoryUserStore) GetUserByUsername(username string) (user.User, error) {
	dao, ok := i.m.Get(username)
	if !ok {
		return user.User{}, ErrUserNotFound
	}

	return MapDAOToUser(dao), nil
}

func (i *InMemoryUserStore) InsertUser(usr user.User) error {
	if i.m.Exists(usr.Username()) {
		return ErrUserAlreadyExists
	}

	i.m.Set(usr.Username(), MapUserToDAO(usr))

	return nil
}

func (i *InMemoryUserStore) UpdateUser(usr user.User) error {
	if !i.m.Exists(usr.Username()) {
		return ErrUserNotFound
	}

	i.m.Set(usr.Username(), MapUserToDAO(usr))
	return nil
}

func (i *InMemoryUserStore) DeleteUser(username string) error {
	if !i.m.Exists(username) {
		return ErrUserNotFound
	}
	i.m.Delete(username)

	return nil
}
