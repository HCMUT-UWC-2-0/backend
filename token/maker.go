package token

import "time"

type Maker interface {
	CreateToken(boi BackOfficerInfo, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
