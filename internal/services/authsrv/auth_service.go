package authsrv

import (
	"errors"
	"github.com/amanbolat/umsystem/internal/auth"
	"github.com/amanbolat/umsystem/internal/services/usersrv"
	"time"
)

var ErrWrongCredentials = errors.New("wrong credentials")
var ErrServerError = errors.New("error on server side")
var ErrTokenExpired = errors.New("auth token expired, please sign in")

type TokenCache interface {
	SaveToken(token auth.AuthToken) error
	GetToken(key string) (auth.AuthToken, error)
	DeleteToken(key string) error
}
type AuthService interface {
	SignIn(username, password string) (authToken string, err error)
	SignOut(tokenKey string) error
	Authorize(tokenKey string) (auth.AuthToken, error)
}

type authService struct {
	tokenCache TokenCache
	tokenTts   time.Duration
	userSrv    usersrv.UserService
}

func NewAuthService(tokenCache TokenCache, usrService usersrv.UserService, tokenTts time.Duration) AuthService {
	return &authService{
		tokenCache: tokenCache,
		tokenTts:   tokenTts,
		userSrv:    usrService,
	}
}

func (a *authService) SignOut(tokenKey string) error {
	return a.tokenCache.DeleteToken(tokenKey)
}

func (a *authService) SignIn(username, password string) (string, error) {
	usr, err := a.userSrv.GetUserByUsername(username)
	if err != nil {
		return "", ErrWrongCredentials
	}

	if !usr.ComparePassword(password) {
		return "", ErrWrongCredentials
	}

	expireAt := time.Now().Add(a.tokenTts)
	token, err := auth.NewToken(username, expireAt)
	if err != nil {
		return "", ErrServerError
	}

	err = a.tokenCache.SaveToken(token)
	if err != nil {
		return "", ErrServerError
	}

	return token.Key, nil
}

func (a *authService) Authorize(tokenKey string) (auth.AuthToken, error) {
	token, err := a.tokenCache.GetToken(tokenKey)
	if err != nil {
		return auth.AuthToken{}, ErrWrongCredentials
	}

	if token.IsExpired() {
		_ = a.tokenCache.DeleteToken(tokenKey)
		return auth.AuthToken{}, ErrTokenExpired
	}

	return token, nil
}
