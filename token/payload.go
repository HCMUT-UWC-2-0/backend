package token

import (
	"errors"
	"time"

	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/google/uuid"
)

type Payload struct {
	ID          uuid.UUID   `json:"id"`
	AccountInfo AccountInfo `json:"student_info"`
	IssuedAt    time.Time   `json:"issued_at"`
	ExpiredAt   time.Time   `json:"expired_at"`
}

// MetaData  interface  `json:"meta_data"`
type AccountInfo struct {
	AccountId string      `json:"account_id"`
	Role      db.RoleType `json:"role"`
}

//type MetaData interface{}

func NewPayload(accountInfo AccountInfo, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenID,
		AccountInfo: accountInfo,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(duration),
	}
	//    MetaData:  metaData,
	return payload, nil
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
