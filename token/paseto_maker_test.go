package token

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/DV-Lab/zuni-backend/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	account := createRandomStudent()
	//	meta_data := MetaData(&userinfo)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(account, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)

	// Get metadatapayload
	b, err := json.Marshal(payload.AccountInfo.AccountId)
	require.NoError(t, err)
	b2, err := json.Marshal(payload.AccountInfo.Role)
	require.NoError(t, err)


	require.NoError(t, err)

	var accountId string
	var role db.RoleType
	err = json.Unmarshal(b, &accountId)
	require.NoError(t, err)
	err = json.Unmarshal(b2, &role)
	require.NoError(t, err)
	require.NoError(t, err)

	fmt.Println(">>>", accountId)
	fmt.Println(">>>", role)
	fmt.Println(">>>", account.AccountId)
	fmt.Println(">>>", account.Role)

	require.Equal(t, account.AccountId, accountId)
	require.Equal(t, account.Role, role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
func createRandomStudent() AccountInfo {
	return AccountInfo{
		AccountId: util.RandomString(6),
		Role: 	db.RoleType(util.RandomRole()),
	}
}
func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(createRandomStudent(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
