package pkg

import (
	"crypto/rand"
	"math/big"
)

const randStringLetters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+"

func String(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(randStringLetters))))
		if err != nil {
			return "", err
		}
		ret[i] = randStringLetters[num.Int64()]
	}

	return string(ret), nil
}
