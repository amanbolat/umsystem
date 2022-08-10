package auth

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	t.Run("new token", func(t *testing.T) {
		username := "user123"
		tts := time.Second * 10
		expireAt := time.Now().Add(tts)
		token, err := NewToken(username, expireAt)
		if err != nil {
			t.Fatal("failed to create new token")
		}

		if token.Key == "" {
			t.Fatal("token key should not be empty")
		}

		if token.Username != username {
			t.Fatalf("token username should be %s", username)
		}

		if token.IsExpired() {
			t.Fatalf("token was just created, it should not be expired")
		}
	})

	t.Run("expired token", func(t *testing.T) {
		tts := time.Millisecond * 500
		expireAt := time.Now().Add(tts)
		token, err := NewToken("user123", expireAt)
		if err != nil {
			t.Fatal("failed to create new token")
		}

		time.Sleep(tts)

		if !token.IsExpired() {
			t.Fatalf("token should be expired by now")
		}
	})

	t.Run("equal tokens", func(t *testing.T) {
		now := time.Now()
		token1, _ := NewToken("user", now)
		token2 := AuthToken{
			Key:      token1.Key,
			Username: token1.Username,
			ExpireAt: token1.ExpireAt,
		}
		if !token1.EqualTo(token2) {
			t.Fatalf("tokens are not equal")
		}
	})
}
