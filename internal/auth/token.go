package auth

import (
	"github.com/amanbolat/umsystem/pkg"
	"time"
)

const defaultTokenLength = 32

type AuthToken struct {
	Key      string
	Username string
	ExpireAt time.Time
}

func (t *AuthToken) EqualTo(t2 AuthToken) bool {
	if t.Key != t2.Key {
		return false
	}

	if t.Username != t2.Username {
		return false
	}

	if !t.ExpireAt.Equal(t2.ExpireAt) {
		return false
	}

	return true
}

func (t *AuthToken) IsExpired() bool {
	return t.ExpireAt.Before(time.Now())
}

func NewToken(username string, expireAt time.Time) (AuthToken, error) {
	randStr, err := pkg.String(defaultTokenLength)
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		Key:      randStr,
		Username: username,
		ExpireAt: expireAt,
	}, nil
}
