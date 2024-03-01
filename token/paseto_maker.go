package token

import (
	"aidanwoods.dev/go-paseto"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.Token
	symmetricKey []byte
}

func CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

}
