package token

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/HCMUT-UWC-2-0/backend/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	bo := createRandomBackOfficer()
	
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(bo, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)

	// Get metadatapayload
	b, err := json.Marshal(payload.BackOfficerInfo.Email)
	require.NoError(t, err)
	b2, err := json.Marshal(payload.BackOfficerInfo.Ssn)
	require.NoError(t, err)


	require.NoError(t, err)

	var email string
	var ssn string
	err = json.Unmarshal(b, &email)
	require.NoError(t, err)
	err = json.Unmarshal(b2, &ssn)
	require.NoError(t, err)
	require.NoError(t, err)

	fmt.Println(">>>", email)
	fmt.Println(">>>", ssn)
	fmt.Println(">>>", bo.Email)
	fmt.Println(">>>", bo.Ssn)

	require.Equal(t, bo.Email, email)
	require.Equal(t, bo.Ssn, ssn)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
func createRandomBackOfficer() BackOfficerInfo {
	return BackOfficerInfo{
		Ssn: util.RandomString(6),
		Email: util.RandomEmail(),
	}
}
func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(createRandomBackOfficer(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
