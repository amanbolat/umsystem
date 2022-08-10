package api

import (
	"bytes"
	"encoding/json"
	"github.com/amanbolat/umsystem/internal/cache"
	"github.com/amanbolat/umsystem/internal/services/authsrv"
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/services/usersrv"
	"github.com/amanbolat/umsystem/internal/store"
	"github.com/amanbolat/umsystem/internal/user"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	roleStore := store.NewInMemoryRoleStore()
	roleSrv := rolesrv.NewRoleService(roleStore)
	usrStore := store.NewInMemoryUserStore()
	usrValidator := user.NewUserValidator(8, 1)
	usrSrv := usersrv.NewUserService(usrStore, roleSrv, usrValidator)
	tokenCache := cache.NewInMemoryTokenCache()
	authSrv := authsrv.NewAuthService(tokenCache, usrSrv, time.Hour*2)
	server := NewServer(usrSrv, authSrv, roleSrv, 9999)

	go func() {
		server.Start()
	}()

	time.Sleep(time.Millisecond * 100)

	username := "username123"
	password := "password123"

	t.Run("create user", func(t *testing.T) {
		req := CreateUserRequest{
			Username: username,
			Password: password,
		}
		b, _ := json.Marshal(&req)
		buf := bytes.NewBuffer(b)

		res, err := http.Post("http://localhost:9999/user", "application/json", buf)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to parse response")
		}

		var createResponse CreateUserResponse
		err = json.Unmarshal(data, &createResponse)
		if err != nil {
			t.Fatalf("failed to unmarshal response")
		}

		if req.Username != createResponse.Username {
			t.Fatalf("usernames should be equal")
		}
	})
}
