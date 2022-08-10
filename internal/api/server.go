package api

import (
	"context"
	"fmt"
	"github.com/amanbolat/umsystem/internal/auth"
	"github.com/amanbolat/umsystem/internal/services/authsrv"
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/services/usersrv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const authTokenHeader = "Token"
const contextTokenKey = "contextTokenKey"

type Server struct {
	usrSrv  usersrv.UserService
	authSrv authsrv.AuthService
	roleSrv rolesrv.RoleService
	e       *echo.Echo
	port    int
}

func NewServer(usrSrv usersrv.UserService, authSrv authsrv.AuthService, roleSrv rolesrv.RoleService, port int) *Server {
	e := echo.New()
	e.HideBanner = true

	s := &Server{
		usrSrv:  usrSrv,
		authSrv: authSrv,
		roleSrv: roleSrv,
		e:       e,
		port:    port,
	}
	s.RegisterHandlers()

	return s
}

func (s *Server) RegisterHandlers() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.POST("/user", s.CreateUser)
	s.e.DELETE("/user/:username", s.DeleteUser)
	s.e.POST("/role", s.CreateRole)
	s.e.DELETE("/role/:roleName", s.DeleteRole)
	s.e.POST("/user/role", s.AddRoleToUser)
	s.e.POST("/signin", s.SignIn)
	s.e.POST("/signout", s.SignOut, s.AuthMiddleware)
	s.e.HEAD("/user/role/:roleName", s.UserHasRole, s.AuthMiddleware)
	s.e.GET("/user/role", s.GetUserRoles, s.AuthMiddleware)
}

func (s *Server) Start() error {
	return s.e.Start(fmt.Sprintf(":%d", s.port))
}

func (s *Server) Stop() error {
	return s.e.Shutdown(context.Background())
}

func (s *Server) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenKey := c.Request().Header.Get(authTokenHeader)
		token, err := s.authSrv.Authorize(tokenKey)
		if err != nil {
			return &echo.HTTPError{
				Code: http.StatusUnauthorized,
			}
		}

		c.Set(contextTokenKey, token)

		return next(c)
	}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
}

func authTokenFromContext(c echo.Context) auth.AuthToken {
	return c.Get(contextTokenKey).(auth.AuthToken)
}

func (s *Server) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	usr, err := s.usrSrv.CreateUser(req.Username, req.Password)
	if err != nil {
		return err
	}

	res := CreateUserResponse{Username: usr.Username()}

	return c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteUser(c echo.Context) error {
	username := c.Param("username")
	err := s.usrSrv.DeleteUser(username)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

type CreateRoleRequest struct {
	RoleName string `json:"role_name"`
}

type CreateRoleResponse struct {
	RoleName string `json:"role_name"`
}

func (s *Server) CreateRole(c echo.Context) error {
	var req CreateRoleRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	role, err := s.roleSrv.CreateRole(req.RoleName)
	if err != nil {
		return err
	}

	res := CreateRoleResponse{RoleName: role.RoleName}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteRole(c echo.Context) error {
	err := s.roleSrv.DeleteRole(c.Param("roleName"))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

type AddRoleToUserRequest struct {
	RoleName string `json:"role_name"`
}

func (s *Server) AddRoleToUser(c echo.Context) error {
	var req AddRoleToUserRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	authToken := authTokenFromContext(c)
	err = s.usrSrv.AssignRole(authToken.Username, req.RoleName)
	if err != nil {
		return err
	}

	return err
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func (s *Server) SignIn(c echo.Context) error {
	var req SignInRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	authToken, err := s.authSrv.SignIn(req.Username, req.Password)
	if err != nil {
		return &echo.HTTPError{
			Code: http.StatusUnauthorized,
		}
	}

	res := SignInResponse{Token: authToken}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) SignOut(c echo.Context) error {
	token := authTokenFromContext(c)
	_ = s.authSrv.SignOut(token.Key)
	return c.NoContent(http.StatusOK)
}

type UserHasRoleRequest struct {
	RoleName string `json:"role_name"`
}

type UserHasRoleResponse struct {
	Result bool `json:"result"`
}

func (s *Server) UserHasRole(c echo.Context) error {
	var req UserHasRoleRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	token := authTokenFromContext(c)

	has, err := s.usrSrv.UserHasRole(token.Username, req.RoleName)
	if err != nil {
		return err
	}
	res := UserHasRoleResponse{Result: has}
	return c.JSON(http.StatusOK, res)
}

type GetUserRolesRequest struct {
	Roles []string
}

func (s *Server) GetUserRoles(c echo.Context) error {
	token := authTokenFromContext(c)

	usr, err := s.usrSrv.GetUserByUsername(token.Username)
	if err != nil {
		return err
	}

	res := GetUserRolesRequest{Roles: usr.Roles()}

	return c.JSON(http.StatusOK, res)
}
