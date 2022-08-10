package authsrv

import (
	"github.com/amanbolat/umsystem/internal/cache"
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/services/usersrv"
	"github.com/amanbolat/umsystem/internal/store"
	"github.com/amanbolat/umsystem/internal/user"
	"testing"
	"time"
)

func TestAuthService(t *testing.T) {
	tokenCache := cache.NewInMemoryTokenCache()
	roleStore := store.NewInMemoryRoleStore()
	roleSrv := rolesrv.NewRoleService(roleStore)
	usrStore := store.NewInMemoryUserStore()
	usrValidator := user.NewUserValidator(3, 3)
	usrSrv := usersrv.NewUserService(usrStore, roleSrv, usrValidator)
	tts := time.Millisecond * 500

	srv := NewAuthService(tokenCache, usrSrv, tts)

	username := "username"
	password := "password"
	_, err := usrSrv.CreateUser(username, password)
	if err != nil {
		t.Fatalf("faield to create user")
	}

	t.Run("sign in - valid", func(t *testing.T) {
		_, err := srv.SignIn(username, password)
		if err != nil {
			t.Fatalf("failed to sign in")
		}
	})

	t.Run("sign in - wrong username", func(t *testing.T) {
		_, err := srv.SignIn("wrong", password)
		if err == nil {
			t.Fatalf("should fail")
		}
	})

	t.Run("sign in - wrong password", func(t *testing.T) {
		_, err := srv.SignIn(username, "wrong")
		if err == nil {
			t.Fatalf("should fail")
		}
	})

	t.Run("authorize - expired token", func(t *testing.T) {
		token, err := srv.SignIn(username, password)
		if err != nil {
			t.Fatalf("failed to sign in")
		}

		time.Sleep(tts)
		_, err = srv.Authorize(token)
		if err == nil {
			t.Fatalf("should fail")
		}
	})

	t.Run("authorize - valid", func(t *testing.T) {
		token, err := srv.SignIn(username, password)
		if err != nil {
			t.Fatalf("failed to sign in")
		}

		_, err = srv.Authorize(token)
		if err != nil {
			t.Fatalf("failed to authorize")
		}
	})

	t.Run("sign out", func(t *testing.T) {
		token, err := srv.SignIn(username, password)
		if err != nil {
			t.Fatalf("failed to sign in")
		}

		_, err = srv.Authorize(token)
		if err != nil {
			t.Fatalf("failed to authorize")
		}

		err = srv.SignOut(token)
		if err != nil {
			t.Fatalf("failed to sign out")
		}
	})
}
