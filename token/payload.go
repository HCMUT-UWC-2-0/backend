package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID              uuid.UUID       `json:"id"`
	BackOfficerInfo BackOfficerInfo `json:"back_officer_info"`
	IssuedAt        time.Time       `json:"issued_at"`
	ExpiredAt       time.Time       `json:"expired_at"`
}

// MetaData  interface  `json:"meta_data"`
type BackOfficerInfo struct {
	Email string `json:"email"`
	Ssn   string `json:"ssn"`
}

//type MetaData interface{}

func NewPayload(boi BackOfficerInfo, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenID,
		BackOfficerInfo: boi,
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
