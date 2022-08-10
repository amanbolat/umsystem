package main

import (
	"context"
	"github.com/amanbolat/umsystem/internal/api"
	"github.com/amanbolat/umsystem/internal/cache"
	"github.com/amanbolat/umsystem/internal/services/authsrv"
	"github.com/amanbolat/umsystem/internal/services/rolesrv"
	"github.com/amanbolat/umsystem/internal/services/usersrv"
	"github.com/amanbolat/umsystem/internal/store"
	"github.com/amanbolat/umsystem/internal/user"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	roleStore := store.NewInMemoryRoleStore()
	roleSrv := rolesrv.NewRoleService(roleStore)
	usrStore := store.NewInMemoryUserStore()
	usrValidator := user.NewUserValidator(8, 1)
	usrSrv := usersrv.NewUserService(usrStore, roleSrv, usrValidator)
	tokenCache := cache.NewInMemoryTokenCache()
	authSrv := authsrv.NewAuthService(tokenCache, usrSrv, time.Hour*2)
	server := api.NewServer(usrSrv, authSrv, roleSrv, 9999)

	errChan := make(chan error, 10)

	go func() {
		errChan <- server.Start()
	}()

	bgCtx := context.Background()
	ctx, stop := signal.NotifyContext(bgCtx, os.Interrupt, os.Kill)
	defer stop()

	log.Println("starting the server")
	select {
	case <-ctx.Done():
		log.Println("stopping the server: ", ctx.Err())
		_ = server.Stop()
		stop()
	case err := <-errChan:
		log.Fatalf("failed to run the server with error: %v", err)
	}
}
