package token

import (
	"testing"
	"time"

	"settl-backend/utils"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := GenerateJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	email := utils.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
