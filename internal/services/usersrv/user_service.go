package usersrv

import (
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/user"
)

type UserStore interface {
	GetUserByUsername(username string) (user.User, error)
	InsertUser(u user.User) error
	UpdateUser(u user.User) error
	DeleteUser(username string) error
}

type UserService interface {
	CreateUser(username, password string) (user.User, error)
	GetUserByUsername(username string) (user.User, error)
	DeleteUser(username string) error
	AssignRole(username string, role string) error
	UserHasRole(username string, role string) (bool, error)
}

type userService struct {
	store         UserStore
	userValidator *user.UserValidator
	roleManager   rolesrv.RoleService
}

func NewUserService(userStore UserStore, roleManager rolesrv.RoleService, userValidator *user.UserValidator) UserService {
	return &userService{
		store:         userStore,
		userValidator: userValidator,
		roleManager:   roleManager,
	}
}

func (u userService) CreateUser(username, password string) (user.User, error) {
	usr := user.NewUser(username, password)
	err := u.userValidator.ValidateUser(usr)
	if err != nil {
		return user.User{}, err
	}
	usr.PreSave()

	err = u.store.InsertUser(usr)
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (u userService) GetUserByUsername(username string) (user.User, error) {
	return u.store.GetUserByUsername(username)
}

func (u userService) DeleteUser(username string) error {
	return u.store.DeleteUser(username)
}

func (u userService) AssignRole(username string, roleName string) error {
	usr, err := u.store.GetUserByUsername(username)
	if err != nil {
		return err
	}

	role, err := u.roleManager.GetRoleByName(roleName)
	if err != nil {
		return err
	}

	usr.AssignRole(role.RoleName)

	return u.store.UpdateUser(usr)
}

func (u userService) UserHasRole(username string, role string) (bool, error) {
	usr, err := u.store.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	return usr.HasRole(role), nil
}
