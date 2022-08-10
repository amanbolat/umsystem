package cache

import (
	"github.com/amanbolat/umsystem/internal/auth"
	"testing"
	"time"
)

func TestInMemoryTokenCache(t *testing.T) {
	cache := NewInMemoryTokenCache()

	token, err := auth.NewToken("user", time.Now())
	if err != nil {
		t.Fatalf("failed to create token")
	}

	t.Run("save token", func(t *testing.T) {
		err := cache.SaveToken(token)
		if err != nil {
			t.Fatalf("failed to save token")
		}
	})

	t.Run("get token", func(t *testing.T) {
		found, err := cache.GetToken(token.Key)
		if err != nil {
			t.Fatalf("failed to findtoken")
		}
		if !found.EqualTo(token) {
			t.Fatalf("found token is not equal to the original one")
		}
	})

	t.Run("delete token", func(t *testing.T) {
		err := cache.DeleteToken(token.Key)
		if err != nil {
			t.Fatalf("failed to delete token")
		}

		_, err = cache.GetToken(token.Key)
		if err == nil {
			t.Fatalf("token was not deleted")
		}
	})
}
