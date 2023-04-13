package token

import "time"

type Maker interface {
	CreateToken(accountInfo AccountInfo, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
